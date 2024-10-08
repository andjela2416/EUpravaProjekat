package data

import (
	"context"
	"dorm-service/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DormRepo struct {
	cli    *mongo.Client
	logger *log.Logger
	client *http.Client
}

func NewDormRepo(ctx context.Context, logger *log.Logger) (*DormRepo, error) {
	dburi := fmt.Sprintf("mongodb://%s:%s/", os.Getenv("DORM_DB_HOST"), os.Getenv("DORM_DB_PORT"))

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
			MaxConnsPerHost:     10,
		},
	}

	// Return repository with logger and DB client
	return &DormRepo{
		logger: logger,
		cli:    client,
		client: httpClient,
	}, nil
}

// Disconnect from database
func (dr *DormRepo) DisconnectMongo(ctx context.Context) error {
	err := dr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (dr *DormRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := dr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		dr.logger.Println(err)
	}

	// Print available databases
	databases, err := dr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		dr.logger.Println(err)
	}
	fmt.Println(databases)
}
func (dr *DormRepo) GetClient() *mongo.Client {
	return dr.cli
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database(os.Getenv("DORM_DB_HOST")).Collection(collectionName)

	return collection
}

func (dr *DormRepo) GetApplication(studentid string, selectionid string) (*models.Application, error) {

	var app models.Application
	selCollection := OpenCollection(dr.cli, "selections")
	studentObjectId, err := primitive.ObjectIDFromHex(studentid)
	selectionObjectid, err := primitive.ObjectIDFromHex(selectionid)

	if err != nil {
		return nil, fmt.Errorf("invalid ID: %v", err)
	}

	var selection models.Selection
	filter := bson.M{"_id": selectionObjectid}
	err = selCollection.FindOne(context.Background(), filter).Decode(&selection)
	if err != nil {
		return nil, fmt.Errorf("selection with ID %s not found", selectionid)
	}

	for _, app := range selection.Applications {
		if app.Student.ID == studentObjectId {
			return app, nil
		}
	}

	return &app, nil
}

func (dr *DormRepo) GetSelection(selectionID string) (*models.Selection, error) {

	var selection models.Selection
	selCollection := OpenCollection(dr.cli, "selections")
	objectId, err := primitive.ObjectIDFromHex(selectionID)

	err = selCollection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&selection)
	if err != nil {
		return nil, fmt.Errorf("no selections not found for id: %s", selectionID)
	}

	return &selection, nil
}

func (dr *DormRepo) InsertSelection(selection *models.Selection) (primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	selCollection := OpenCollection(dr.cli, "selections")
	selection.Id = primitive.NewObjectID()
	result, err := selCollection.InsertOne(ctx, &selection)
	if err != nil {
		dr.logger.Println(err)
		return primitive.NewObjectID(), err
	}
	dr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return selection.Id, nil
}

func (dr *DormRepo) UpdateSelection(selectionID string, selection *models.Selection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	selCollection := OpenCollection(dr.cli, "selections")
	objectId, err := primitive.ObjectIDFromHex(selectionID)
	if err != nil {
		return fmt.Errorf("invalid ID: %v", err)
	}

	updateData := bson.M{
		"$set": bson.M{
			"startdate":    selection.StartDate,
			"enddate":      selection.EndDate,
			"buildingId":   selection.BuildingId,
			"applications": selection.Applications,
		},
	}

	result, err := selCollection.UpdateOne(ctx, bson.M{"_id": objectId}, updateData)
	if err != nil {
		return fmt.Errorf("could not update selection with id: %s, error: %v", selectionID, err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no selection found with id: %s", selectionID)
	}

	dr.logger.Printf("Updated selection with id: %s", selectionID)
	return nil
}

func (dr *DormRepo) DeleteSelection(selectionID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	selCollection := OpenCollection(dr.cli, "selections")
	objectId, err := primitive.ObjectIDFromHex(selectionID)
	if err != nil {
		return fmt.Errorf("invalid ID: %v", err)
	}

	result, err := selCollection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return fmt.Errorf("could not delete selection with id: %s, error: %v", selectionID, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no selection found with id: %s", selectionID)
	}

	dr.logger.Printf("Deleted selection with id: %s", selectionID)
	return nil
}

func (dr *DormRepo) CheckSelectionOverlap(selection models.Selection, isUpdate bool) error {
	selectionCollection := OpenCollection(dr.cli, "selections")

	currentStartDate, err := time.Parse("02-01-2006", selection.StartDate)
	if err != nil {
		return fmt.Errorf("invalid start date format: %v", err)
	}
	currentEndDate, err := time.Parse("02-01-2006", selection.EndDate)
	if err != nil {
		return fmt.Errorf("invalid end date format: %v", err)
	}

	now := time.Now()

	filter := bson.M{
		"startdate":  bson.M{"$gte": now.Format("02-01-2006")},
		"buildingId": selection.BuildingId,
	}
	if isUpdate {
		filter = bson.M{
			"startdate":  bson.M{"$gte": now.Format("02-01-2006")},
			"buildingId": selection.BuildingId,
			"_id":        bson.M{"$ne": selection.Id},
		}
	}
	cursor, err := selectionCollection.Find(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error querying selections: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var otherSelection models.Selection
		if err := cursor.Decode(&otherSelection); err != nil {
			return fmt.Errorf("error decoding selection: %v", err)
		}

		otherStartDate, err := time.Parse("02-01-2006", otherSelection.StartDate)
		if err != nil {
			return fmt.Errorf("invalid start date format in other selection: %v", err)
		}
		otherEndDate, err := time.Parse("02-01-2006", otherSelection.EndDate)
		if err != nil {
			return fmt.Errorf("invalid end date format in other selection: %v", err)
		}

		if (currentStartDate.Before(otherEndDate) && currentEndDate.After(otherStartDate)) ||
			(otherStartDate.Before(currentEndDate) && otherEndDate.After(currentStartDate) ||
				currentStartDate.Equal(otherStartDate) || currentStartDate.Equal(otherEndDate) ||
				currentEndDate.Equal(otherStartDate) || currentEndDate.Equal(otherEndDate)) {

			return fmt.Errorf("the date you selected overlaps with already existing ones: %s - %s", otherStartDate.String(), otherEndDate.String())
		}
	}

	return nil
}

func (dr *DormRepo) Insertapplications(application *models.Application) error {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	application.Status = "Pending"
	appsCollection := OpenCollection(dr.cli, "applications")
	result, err := appsCollection.InsertOne(ctx, &application)
	if err != nil {
		dr.logger.Println(err)
		return err
	}
	dr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (dr *DormRepo) InsertApp(app models.Application, selectionId string) (models.Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	var errorApp models.Application

	selCollection := OpenCollection(dr.cli, "selections")

	buildingObjectId, err := primitive.ObjectIDFromHex(selectionId)
	if err != nil {
		return errorApp, fmt.Errorf("invalid building ID: %v", err)
	}

	selection, err := dr.GetSelection(selectionId)
	if err != nil {
		return errorApp, fmt.Errorf("failed to get building: %v", err)
	}
	app.Id = primitive.NewObjectID()
	app.Status = "Pending"

	selection.Applications = append(selection.Applications, &app)

	_, err = selCollection.UpdateOne(
		ctx,
		bson.M{"_id": buildingObjectId},
		bson.M{"$set": bson.M{"applications": selection.Applications}},
	)
	if err != nil {
		return errorApp, fmt.Errorf("error updating building: %v", err)
	}

	return app, nil
}

func (dr *DormRepo) DeleteApp(appUserID string, selectionId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	selCollection := OpenCollection(dr.cli, "selections")

	selectionObjectId, err := primitive.ObjectIDFromHex(selectionId)
	if err != nil {
		return fmt.Errorf("invalid selection ID: %v", err)
	}

	selection, err := dr.GetSelection(selectionId)
	if err != nil {
		return fmt.Errorf("failed to get selection: %v", err)
	}

	// Find and remove the application with the given user ID
	var updatedApplications []*models.Application
	appFound := false
	for _, app := range selection.Applications {
		if app.Student != nil && app.Student.ID.Hex() == appUserID {
			appFound = true
		} else {
			updatedApplications = append(updatedApplications, app)
		}
	}

	if !appFound {
		return fmt.Errorf("no application found with user ID: %s", appUserID)
	}

	// Update the selection with the modified applications
	_, err = selCollection.UpdateOne(
		ctx,
		bson.M{"_id": selectionObjectId},
		bson.M{"$set": bson.M{"applications": updatedApplications}},
	)
	if err != nil {
		return fmt.Errorf("error updating selection: %v", err)
	}

	return nil
}

func (dr *DormRepo) GetAllApplications() ([]*models.Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	appsCollection := OpenCollection(dr.cli, "applications")

	var apps []*models.Application
	appCursor, err := appsCollection.Find(ctx, bson.M{})
	if err != nil {
		dr.logger.Println(err)
		return nil, err
	}
	if err = appCursor.All(ctx, &apps); err != nil {
		dr.logger.Println(err)
		return nil, err
	}
	return apps, nil
}

func (dr *DormRepo) InsertBuilding(building models.Building) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	building.Id = primitive.NewObjectID()
	buildingCollection := OpenCollection(dr.cli, "buildings")
	result, err := buildingCollection.InsertOne(ctx, &building)
	if err != nil {
		dr.logger.Println(err)
		return building.Id, err
	}
	dr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return building.Id, nil

}

func (dr *DormRepo) GetBuilding(id string) (*models.Building, error) {

	var building models.Building
	buildingCollection := OpenCollection(dr.cli, "buildings")
	buildingObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid building ID: %v", err)
	}

	err = buildingCollection.FindOne(context.Background(), bson.M{"_id": buildingObjectID}).Decode(&building)
	if err != nil {
		return nil, err
	}

	return &building, nil
}
func (dr *DormRepo) DeleteBuilding(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	buildingObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid building ID: %v", err)
	}

	buildingCollection := OpenCollection(dr.cli, "buildings")

	result, err := buildingCollection.DeleteOne(ctx, bson.M{"_id": buildingObjectID})
	if err != nil {
		return fmt.Errorf("error deleting building: %v", err)
	}

	if result.DeletedCount != 1 {
		return fmt.Errorf("building not found")
	}

	return nil
}
func (dr *DormRepo) UpdateBuilding(buildingID primitive.ObjectID, address string, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	buildingCollection := OpenCollection(dr.cli, "buildings")

	update := bson.M{}
	if address != "" {
		update["address"] = address
		update["name"] = name
	}

	_, err := buildingCollection.UpdateOne(
		ctx,
		bson.M{"_id": buildingID},
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating building: %v", err)
	}
	return nil
}
func (dr *DormRepo) InsertRoom(insertedCapacity int, insertedBuildingId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	buildingCollection := OpenCollection(dr.cli, "buildings")
	var room models.Room

	buildingObjectId, err := primitive.ObjectIDFromHex(insertedBuildingId)
	if err != nil {
		return fmt.Errorf("invalid building ID: %v", err)
	}

	building, err := dr.GetBuilding(insertedBuildingId)
	if err != nil {
		return fmt.Errorf("failed to get building: %v", err)
	}

	room.Room_Number = len(building.Rooms) + 1
	room.Building_Id = buildingObjectId
	room.Capacity = insertedCapacity

	building.Rooms = append(building.Rooms, &room)

	_, err = buildingCollection.UpdateOne(
		ctx,
		bson.M{"_id": buildingObjectId},
		bson.M{"$set": bson.M{"rooms": building.Rooms}},
	)
	if err != nil {
		return fmt.Errorf("error updating building: %v", err)
	}

	return nil
}

// EditRoom updates the details of a room in a building
func (dr *DormRepo) EditRoom(roomNumber int, buildingId primitive.ObjectID, updatedRoom *models.Room) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	roomCollection := OpenCollection(dr.cli, "rooms")

	// Define the filter to find the specific room based on room_number and building_id
	filter := bson.M{
		"room_number": roomNumber,
		"building_id": buildingId,
	}

	// Define the fields to update
	update := bson.M{
		"$set": bson.M{
			"room_number": updatedRoom.Room_Number,
			"capacity":    updatedRoom.Capacity,
			"students":    updatedRoom.Students, // Update the students list if needed
		},
	}

	// Perform the update operation
	result, err := roomCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update room: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no room found with room number %d in building", roomNumber)
	}

	dr.logger.Printf("Room %d in building %s updated successfully", roomNumber, buildingId.Hex())
	return nil
}
func (dr *DormRepo) GetRoom(number int, buildingId string) (*models.Room, error) {
	buildingCollection := OpenCollection(dr.cli, "buildings")

	buildingObjectId, err := primitive.ObjectIDFromHex(buildingId)
	if err != nil {
		return nil, fmt.Errorf("invalid building ID: %v", err)
	}

	var building models.Building
	filter := bson.M{"_id": buildingObjectId}
	err = buildingCollection.FindOne(context.Background(), filter).Decode(&building)
	if err != nil {
		return nil, fmt.Errorf("building with ID %s not found", buildingId)
	}

	for _, room := range building.Rooms {
		if room.Room_Number == number {
			return room, nil
		}
	}

	return nil, fmt.Errorf("room #%d in building %s not found", number, buildingId)
}
