package data

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type FoodServiceRepo struct {
	cli    *mongo.Client
	logger *log.Logger
	client *http.Client
	store  *sessions.CookieStore
}

func NewFoodServiceRepo(ctx context.Context, logger *log.Logger) (*FoodServiceRepo, error) {
	dburi := fmt.Sprintf("mongodb://%s:%s/", os.Getenv("FOOD_DB_HOST"), os.Getenv("FOOD_DB_PORT"))

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	store := sessions.NewCookieStore([]byte("super-secret-key"))
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
			MaxConnsPerHost:     10,
		},
	}

	// Return repository with logger and DB client
	return &FoodServiceRepo{
		logger: logger,
		cli:    client,
		client: httpClient,
		store:  store,
	}, nil
}

// Disconnect from database
func (pr *FoodServiceRepo) DisconnectMongo(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (rr *FoodServiceRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := rr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		rr.logger.Println(err)
	}

	// Print available databases
	databases, err := rr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		rr.logger.Println(err)
	}
	fmt.Println(databases)
}

// GetAllFoodOfStudents
func (rr *FoodServiceRepo) GetAllFoodOfStudents() (*Students, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	studentsCollection := rr.getCollection("students")

	var students Students
	studentCursor, err := studentsCollection.Find(ctx, bson.M{})
	if err != nil {
		rr.logger.Println(err)
		return nil, err
	}
	if err = studentCursor.All(ctx, &students); err != nil {
		rr.logger.Println(err)
		return nil, err
	}
	return &students, nil
}

func (rr *FoodServiceRepo) GetTokenFromSession(r *http.Request) (string, error) {
	session, err := rr.store.Get(r, "session-name")
	if err != nil {
		return "", err
	}

	token, ok := session.Values["token"].(string)
	if !ok {
		return "", errors.New("token not found in session")
	}

	return token, nil
}

func (rr *FoodServiceRepo) GetLoggedUser(r *http.Request) (*AuthUser, error) {
	token, err := rr.GetTokenFromSession(r)
	if err != nil {
		return nil, err
	}

	meEndpoint := "http://localhost:8080/user/me"

	req, err := http.NewRequest("GET", meEndpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
		return nil, errors.New("unexpected status code")
	}

	var user AuthUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, err
	}

	return &user, nil
}
func (rr *FoodServiceRepo) CreateFoodEntry(r *http.Request, foodData *Food) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Postavi dummy userID jer trenutno ne koristiš ulogovanog korisnika
	dummyUserID := primitive.NilObjectID
	foodData.UserID = dummyUserID
	foodData.Stanje = Neporucena

	// Loguj podatke pre umetanja
	fmt.Printf("Inserting food data: %+v\n", foodData)

	foodCollection := rr.getCollection("food")

	// Umetanje u MongoDB
	_, err := foodCollection.InsertOne(ctx, foodData)
	if err != nil {
		fmt.Println("Error inserting food data:", err) // Loguj grešku umetanja
		return err
	}

	return nil
}

// editFood
func (rr *FoodServiceRepo) EditFoodForStudent(studentID string, newFood string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	studentsCollection := rr.getCollection("students")

	filter := bson.M{"student_id": studentID}
	update := bson.M{"$set": bson.M{"food": newFood}}

	_, err := studentsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		rr.logger.Println(err)
		return err
	}
	rr.logger.Printf("Food updated successfully for student with ID: %s\n", studentID)
	return nil
}

var therapiesList Therapies

func CacheTherapies(therapies Therapies) {
	therapiesList = append(therapiesList, therapies...)
}

func GetCachedTherapies() Therapies {
	return therapiesList
}

// funkcija dobavlja sve terapije iz Food servisa.
func (rr *FoodServiceRepo) GetAllTherapiesFromFoodService() (Therapies, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	therapiesCollection := rr.cli.Database("MongoDatabase").Collection("therapies")

	var therapies Therapies
	cursor, err := therapiesCollection.Find(ctx, bson.M{})
	if err != nil {
		rr.logger.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &therapies); err != nil {
		rr.logger.Println(err)
		return nil, err
	}

	return therapies, nil
}

func (rr *FoodServiceRepo) SaveTherapyData(therapyData *TherapyData) error {

	therapiesList = append(therapiesList, therapyData)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	therapiesCollection := rr.getCollection("therapies")

	// Insert therapy data into therapies collection
	_, err := therapiesCollection.InsertOne(ctx, therapyData)
	if err != nil {
		rr.logger.Println(err)
		return err
	}
	return nil
}

func (rr *FoodServiceRepo) ClearTherapiesCache() error {
	therapiesList = Therapies{}
	return nil
}

func (rr *FoodServiceRepo) UpdateTherapyStatus(therapyID primitive.ObjectID, status Status) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	therapiesCollection := rr.getCollection("therapies")

	filter := bson.M{"therapyId": therapyID}
	update := bson.M{"$set": bson.M{"status": status}}
	result, err := therapiesCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("therapy with ID %s not found in database", therapyID.Hex())
	}

	updatedTherapy := &TherapyData{
		ID:     therapyID,
		Status: status,
	}
	if err := rr.SendTherapyDataToHealthCareService(updatedTherapy); err != nil {
		return err
	}

	return nil
}

func (rr *FoodServiceRepo) SendTherapyDataToHealthCareService(therapy *TherapyData) error {

	therapyJSON, err := json.Marshal(therapy)
	if err != nil {
		rr.logger.Println("Error serializing therapy data:", err)
		return err
	}

	healthCareHost := os.Getenv("HEALTHCARE_SERVICE_HOST")
	healthCarePort := os.Getenv("HEALTHCARE_SERVICE_PORT")
	healthCareEndpoint := fmt.Sprintf("http://%s:%s/updateTherapy", healthCareHost, healthCarePort)

	req, err := http.NewRequest("PUT", healthCareEndpoint, bytes.NewBuffer(therapyJSON))
	if err != nil {
		rr.logger.Println("Error creating request to health care service:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := rr.client.Do(req)
	if err != nil {
		rr.logger.Println("Error sending request to health care service:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		rr.logger.Println("Health care service returned non-OK status code:", resp.StatusCode)
		return errors.New("health care service returned non-OK status code")
	}

	return nil
}

// GetAllTherapiesFromHealthCareService funkcija dobavlja sve terapije iz HealthCare servisa.
func (rr *FoodServiceRepo) GetAllTherapiesFromHealthCareService() (Therapies, error) {
	healthCareHost := os.Getenv("HEALTHCARE_SERVICE_HOST")
	healthCarePort := os.Getenv("HEALTHCARE_SERVICE_PORT")
	healthCareEndpoint := fmt.Sprintf("http://%s:%s/therapies", healthCareHost, healthCarePort)

	req, err := http.NewRequest("GET", healthCareEndpoint, nil)
	if err != nil {
		rr.logger.Println("Error creating request to health care service:", err)
		return nil, err
	}

	resp, err := rr.client.Do(req)
	if err != nil {
		rr.logger.Println("Error sending request to health care service:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		rr.logger.Println("Health care service returned non-OK status code:", resp.StatusCode)
		return nil, fmt.Errorf("health care service returned non-OK status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rr.logger.Println("Error reading response from health care service:", err)
		return nil, err
	}

	var therapies Therapies
	if err := json.Unmarshal(body, &therapies); err != nil {
		rr.logger.Println("Error parsing response from health care service:", err)
		return nil, err
	}

	CacheTherapies(therapies)

	return therapies, nil
}

func (rr *FoodServiceRepo) getCollection(collectionName string) *mongo.Collection {
	return rr.cli.Database("MongoDatabase").Collection(collectionName)
}
