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

type HealthCareRepo struct {
	cli          *mongo.Client
	logger       *log.Logger
	client       *http.Client
	allTherapies Therapies
}

func NewHealthCareRepo(ctx context.Context, logger *log.Logger) (*HealthCareRepo, error) {
	dburi := fmt.Sprintf("mongodb://%s:%s/", os.Getenv("HEALTHCARE_DB_HOST"), os.Getenv("HEALTHCARE_DB_PORT"))

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
	return &HealthCareRepo{
		logger: logger,
		cli:    client,
		client: httpClient,
	}, nil
}

// Disconnect from database
func (pr *HealthCareRepo) DisconnectMongo(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (rr *HealthCareRepo) Ping() {
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

// mongo
func (rr *HealthCareRepo) InsertStudent(student *Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	studentsCollection := rr.getCollection("students")

	result, err := studentsCollection.InsertOne(ctx, &student)
	if err != nil {
		rr.logger.Println(err)
		return err
	}
	rr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (rr *HealthCareRepo) GetAllStudents() (*Students, error) {
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

// ScheduleAppointment zakazuje pregled za određenog studenta.
func (rr *HealthCareRepo) ScheduleAppointment(appointmentData *AppointmentData) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	examinationsCollection := rr.getCollection("examinations")

	_, err := examinationsCollection.InsertOne(ctx, appointmentData)
	if err != nil {
		return err
	}

	return nil
}

// GetAllAppointments vraća sve zakazane termine pregleda.
func (rr *HealthCareRepo) GetAllAppointments() (*Appointments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	examinationsCollection := rr.getCollection("examinations")
	// Pronađi sve zakazane termine pregleda
	cursor, err := examinationsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var appointments Appointments
	studentCursor, err := examinationsCollection.Find(ctx, bson.M{})
	if err != nil {
		rr.logger.Println(err)
		return nil, err
	}
	if err = studentCursor.All(ctx, &appointments); err != nil {
		rr.logger.Println(err)
		return nil, err
	}
	return &appointments, nil
}

// SaveTherapyData funkcija čuva podatke o terapiji u bazi podataka
func (rr *HealthCareRepo) SaveTherapyData(therapyData *TherapyData) error {
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

// GetAllTherapies dohvata sve terapije iz baze podataka.
func (rr *HealthCareRepo) GetAllTherapies() (Therapies, error) {
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

// ShareTherapyDataWithDietService funkcija deli informacije o terapijama sa službom ishrane.
func (rr *HealthCareRepo) SaveAndShareTherapyDataWithDietService(therapyData *TherapyData) error {
	// 1. Zapise podatke o terapiji
	if err := rr.SaveTherapyData(therapyData); err != nil {
		return err
	}

	// Dodaj terapiju u listu svih terapija
	rr.allTherapies = append(rr.allTherapies, therapyData)

	// 2. Pošalje podatke o terapiji službi ishrane
	if err := rr.SendTherapyDataToDietService(therapyData); err != nil {
		return err
	}

	return nil
}

// SendTherapyDataToDietService funkcija šalje podatke o terapiji službi ishrane
func (rr *HealthCareRepo) SendTherapyDataToDietService(therapyData *TherapyData) error {
	/*
		therapyJSON, err := json.Marshal(therapyData)
		if err != nil {
			rr.logger.Println("Error serializing therapy data:", err)
			return err
		}

		foodServiceHost := os.Getenv("FOOD_SERVICE_HOST")
		foodServicePort := os.Getenv("FOOD_SERVICE_PORT")
		foodServiceEndpoint := fmt.Sprintf("http://%s:%s/therapy", foodServiceHost, foodServicePort)

		req, err := http.NewRequest("POST", foodServiceEndpoint, bytes.NewBuffer(therapyJSON))
		if err != nil {
			rr.logger.Println("Error creating request to food service:", err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		// Šaljemo zahtev servisu ishrane
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			rr.logger.Println("Error sending request to food service:", err)
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			rr.logger.Println("Food service returned non-OK status code:", resp.StatusCode)
			return errors.New("food service returned non-OK status code")
		}*/

	return nil
}

func (rr *HealthCareRepo) getCollection(collectionName string) *mongo.Collection {
	return rr.cli.Database("MongoDatabase").Collection(collectionName)
}
