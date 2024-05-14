package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	"time"
)

type FoodServiceRepo struct {
	cli    *mongo.Client
	logger *log.Logger
	client *http.Client
}

func NewFoodServiceRepo(ctx context.Context, logger *log.Logger) (*FoodServiceRepo, error) {
	dburi := fmt.Sprintf("mongodb://%s:%s/", os.Getenv("FOODSERVICE_DB_HOST"), os.Getenv("FOODSERVICE_DB_PORT"))

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
	return &FoodServiceRepo{
		logger: logger,
		cli:    client,
		client: httpClient,
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

	studentsCollection := rr.getCollection()

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

// editFood
func (rr *FoodServiceRepo) EditFoodForStudent(studentID string, newFood string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	studentsCollection := rr.getCollection()

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

func (rr *FoodServiceRepo) getCollection() *mongo.Collection {
	appointmentDatabase := rr.cli.Database("MongoDatabase")
	appointmentsCollection := appointmentDatabase.Collection("students")
	return appointmentsCollection
}
