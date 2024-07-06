package repositories

import (
	"context"
	"log"
	"os"
	"time"

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

func CreateNotification(client *mongo.Client, notification *Notification) error {
	collection := client.Database("universityDB").Collection("notifications")
	notification.ID = primitive.NewObjectID()
	notification.CreatedAt = time.Now()
	_, err := collection.InsertOne(context.Background(), notification)
	return err
}

func GetNotificationByID(client *mongo.Client, id string) (*Notification, error) {
	collection := client.Database("universityDB").Collection("notifications")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var notification Notification
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&notification)
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func GetAllNotifications(client *mongo.Client) (Notifications, error) {
	collection := client.Database("universityDB").Collection("notifications")

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var notifications Notifications
	for cur.Next(context.Background()) {
		var notification Notification
		err := cur.Decode(&notification)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func DeleteNotification(client *mongo.Client, id string) error {
	collection := client.Database("universityDB").Collection("notifications")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}

func (ar *StudentRepo) getCollection() *mongo.Collection {
	accommodationDatabase := ar.cli.Database("universityDB")
	accommodationsCollection := accommodationDatabase.Collection("student")
	return accommodationsCollection
}
