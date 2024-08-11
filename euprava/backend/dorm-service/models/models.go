package models

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name  *string            `json:"last_name" validate:"required,min=2,max=100"`
	Email      *string            `json:"email" validate:"email,required"`
	Password   *string            `json:"password" validate:"required,min=8"`
	Phone      *string            `json:"phone" validate:"required"`
	Address    *string            `json:"address" validate:"required"`
}

type Student struct {
	Uid         string     `json:"uid"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	FinanceType string     `json:"finance_type"`
	StudyInfo   *StudyInfo `json:"study_info"`
}

type StudyInfo struct {
	HighschoolGPA float64 `json:"highschool_gpa"`
	GPA           float64 `json:"gpa"`
	ESBP          int     `json:"esbp"`
	Year          int     `json:"year"`
}

type Application struct {
	Status string `json:"status"` //accepted / rejected / pending
	User   *User  `json:"user"`
}

type Selection struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Building struct {
	Id      primitive.ObjectID `bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Address string             `json:"address" bson:"address"`
	Rooms   Rooms              `json:"rooms,omitempty" bson:"rooms"`
}

type Room struct {
	Room_Number int                `json:"room_number,omitempty" bson:"room_number"`
	Capacity    int                `json:"capacity" bson:"capacity"`
	Building_Id primitive.ObjectID `json:"building_id" bson:"building_id"`
	Students    *Students          `json:"students,omitempty" bson:"students,omitempty"`
}

type Students []*Student
type Rooms []*Room

func (o *Students) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Rooms) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}
