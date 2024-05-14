package data

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (dr *DormRepo) Insertapplications(application *Application) error {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	appsCollection := OpenCollection(dr.cli, "applications")
	result, err := appsCollection.InsertOne(ctx, &application)
	if err != nil {
		dr.logger.Println(err)
		return err
	}
	dr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (dr *DormRepo) GetAllapplications() (*Rooms, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	appsCollection := OpenCollection(dr.cli, "applications")

	var rooms Rooms
	roomCursor, err := appsCollection.Find(ctx, bson.M{})
	if err != nil {
		dr.logger.Println(err)
		return nil, err
	}
	if err = roomCursor.All(ctx, &rooms); err != nil {
		dr.logger.Println(err)
		return nil, err
	}
	return &rooms, nil
}

func (dr *DormRepo) getCollection() *mongo.Collection {
	appointmentDatabase := dr.cli.Database("MongoDatabase")
	appointmentsCollection := appointmentDatabase.Collection("rooms")
	return appointmentsCollection
}
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database(os.Getenv("DORM_DB_HOST")).Collection(collectionName)

	return collection
}
