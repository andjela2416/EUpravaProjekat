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

type Repository struct {
	cli    *mongo.Client
	logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*Repository, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &Repository{
		cli:    client,
		logger: logger,
	}, nil
}

func (r *Repository) Disconnect(ctx context.Context) error {
	err := r.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getCollection(collectionName string) *mongo.Collection {
	db := r.cli.Database("universityDB")
	return db.Collection(collectionName)
}

func (r *Repository) CreateStudent(student *Student) error {
	collection := r.getCollection("student")
	result, err := collection.InsertOne(context.TODO(), student)
	if err != nil {
		return err
	}
	student.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *Repository) GetStudentByID(userID string) (*Student, error) {
	collection := r.getCollection("student")
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	var student Student
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&student)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *Repository) UpdateStudent(student *Student) error {
	collection := r.getCollection("student")
	_, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": student.ID}, student)
	return err
}

func (r *Repository) DeleteStudent(userID string) error {
	collection := r.getCollection("student")
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}

func (r *Repository) CreateUniversity(university *University) error {
	collection := r.getCollection("university")
	result, err := collection.InsertOne(context.TODO(), university)
	if err != nil {
		return err
	}
	university.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *Repository) GetUniversityByID(universityID string) (*University, error) {
	collection := r.getCollection("university")
	objectID, err := primitive.ObjectIDFromHex(universityID)
	if err != nil {
		return nil, err
	}
	var university University
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&university)
	if err != nil {
		return nil, err
	}
	return &university, nil
}

func (r *Repository) UpdateUniversity(university *University) error {
	collection := r.getCollection("university")
	_, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": university.ID}, university)
	return err
}

func (r *Repository) DeleteUniversity(universityID string) error {
	collection := r.getCollection("university")
	objectID, err := primitive.ObjectIDFromHex(universityID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}

func (r *Repository) CreateDepartment(department *Department) error {
	collection := r.getCollection("department")
	result, err := collection.InsertOne(context.TODO(), department)
	if err != nil {
		return err
	}
	department.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *Repository) GetDepartmentByID(departmentID string) (*Department, error) {
	collection := r.getCollection("department")
	objectID, err := primitive.ObjectIDFromHex(departmentID)
	if err != nil {
		return nil, err
	}
	var department Department
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&department)
	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *Repository) UpdateDepartment(department *Department) error {
	collection := r.getCollection("department")
	_, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": department.ID}, department)
	return err
}

func (r *Repository) DeleteDepartment(departmentID string) error {
	collection := r.getCollection("department")
	objectID, err := primitive.ObjectIDFromHex(departmentID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}

func (r *Repository) CreateProfessor(professor *Professor) error {
	collection := r.getCollection("professor")
	result, err := collection.InsertOne(context.TODO(), professor)
	if err != nil {
		return err
	}
	professor.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *Repository) GetProfessorByID(professorID string) (*Professor, error) {
	collection := r.getCollection("professor")
	objectID, err := primitive.ObjectIDFromHex(professorID)
	if err != nil {
		return nil, err
	}
	var professor Professor
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&professor)
	if err != nil {
		return nil, err
	}
	return &professor, nil
}

func (r *Repository) UpdateProfessor(professor *Professor) error {
	collection := r.getCollection("professor")
	_, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": professor.ID}, professor)
	return err
}

func (r *Repository) DeleteProfessor(professorID string) error {
	collection := r.getCollection("professor")
	objectID, err := primitive.ObjectIDFromHex(professorID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}
