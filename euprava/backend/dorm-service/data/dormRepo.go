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

func (dr *DormRepo) GetApplication(studentid string) (*models.Application, error) {

	var app models.Application
	appsCollection := OpenCollection(dr.cli, "applications")

	err := appsCollection.FindOne(context.Background(), bson.M{"student.student_id": studentid}).Decode(&app)
	if err != nil {
		return nil, fmt.Errorf("No applications not found for student id: %s", studentid)
	}

	return &app, nil
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

func (dr *DormRepo) GetAllapplications() (*models.Application, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	appsCollection := OpenCollection(dr.cli, "applications")

	var apps models.Application
	roomCursor, err := appsCollection.Find(ctx, bson.M{})
	if err != nil {
		dr.logger.Println(err)
		return nil, err
	}
	if err = roomCursor.All(ctx, &apps); err != nil {
		dr.logger.Println(err)
		return nil, err
	}
	return &apps, nil
}

func (dr *DormRepo) InsertBuilding(building models.Building) (error, primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	building.Id = primitive.NewObjectID()
	buildingCollection := OpenCollection(dr.cli, "buildings")
	result, err := buildingCollection.InsertOne(ctx, &building)
	if err != nil {
		dr.logger.Println(err)
		return err, building.Id
	}
	dr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil, building.Id

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

	// Convert string ID to ObjectID
	buildingObjectId, err := primitive.ObjectIDFromHex(insertedBuildingId)
	if err != nil {
		return fmt.Errorf("invalid building ID: %v", err)
	}

	// Get the building
	building, err := dr.GetBuilding(insertedBuildingId)
	if err != nil {
		return fmt.Errorf("failed to get building: %v", err)
	}

	// Set room properties
	room.Room_Number = len(building.Rooms) + 1
	room.Building_Id = buildingObjectId
	room.Capacity = insertedCapacity

	// Update the building's rooms slice
	building.Rooms = append(building.Rooms, &room)

	// Update the building document in MongoDB
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

func (dr *DormRepo) GetRoom(number int, buildingId string) (*models.Room, error) {
	var room models.Room
	buildingCollection := OpenCollection(dr.cli, "buildings")
	filter := bson.M{
		"rooms.room_number": number,
		"_id":               buildingId,
	}

	err := buildingCollection.FindOne(context.Background(), filter).Decode(&room)
	if err != nil {
		return nil, fmt.Errorf("room #%d in building %s not found", number, buildingId)
	}

	return &room, nil
}
