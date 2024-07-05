package repositories

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*StudentRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &StudentRepo{
		cli:    client,
		logger: logger,
	}, nil
}

func (ar *StudentRepo) Disconnect(ctx context.Context) error {
	err := ar.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func Createstudent(client *mongo.Client, student *Student) error {
	studentCollection := client.Database("universityDB").Collection("student")

	result, err := studentCollection.InsertOne(context.TODO(), student)
	if err != nil {
		return err
	}

	student.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func GetstudentByID(client *mongo.Client, userID string) (*Student, error) {
	studentCollection := client.Database("universityDB").Collection("student")

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	var student Student
	err = studentCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (ar *StudentRepo) getCollection() *mongo.Collection {
	accommodationDatabase := ar.cli.Database("universityDB")
	accommodationsCollection := accommodationDatabase.Collection("student")
	return accommodationsCollection
}
