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

type UniRepo struct {
	cli    *mongo.Client
	logger *log.Logger
	client *http.Client
}

func NewUniRepo(ctx context.Context, logger *log.Logger) (*UniRepo, error) {
	dburi := fmt.Sprintf("mongodb://%s:%s/", os.Getenv("UNIVERSITY_DB_HOST"), os.Getenv("UNIVERSITY_DB_PORT"))

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
	return &UniRepo{
		logger: logger,
		cli:    client,
		client: httpClient,
	}, nil
}

// Disconnect from database
func (ur *UniRepo) DisconnectMongo(ctx context.Context) error {
	err := ur.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (ur *UniRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := ur.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		ur.logger.Println(err)
	}

	// Print available databases
	databases, err := ur.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		ur.logger.Println(err)
	}
	fmt.Println(databases)
}

func (dr *UniRepo) InsertStudent(student *Student) error {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	studCollection := OpenCollection(dr.cli, "students")
	result, err := studCollection.InsertOne(ctx, &student)
	if err != nil {
		dr.logger.Println(err)
		return err
	}
	dr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (dr *UniRepo) GetStudentByStudentID(id string) (*Student, error) {

	studCollection := OpenCollection(dr.cli, "students")
	var student *Student

	err := studCollection.FindOne(context.Background(), bson.M{"student_id": id}).Decode(&student)
	if err != nil {
		dr.logger.Println(err.Error())
		return nil, err
	}
	return student, nil
}
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database(os.Getenv("UNIVERSITY_DB_HOST")).Collection(collectionName)

	return collection
}
