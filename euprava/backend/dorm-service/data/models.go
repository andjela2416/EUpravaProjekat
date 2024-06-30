package data

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	StudentId   string     `json:"student_id"`
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

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name  *string            `json:"last_name" validate:"required,min=2,max=100"`
	Email      *string            `json:"email" validate:"email,required"`
	Password   *string            `json:"password" validate:"required,min=8"`
	Phone      *string            `json:"phone" validate:"required"`
	Address    *string            `json:"address" validate:"required"`
}

type Application struct {
	Status string `json:"status"` //accepted / rejected / pending
	User   *User  `json:"user"`
}
type Selection struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Room struct {
	RoomNumber int       `json:"room_number"`
	Capacity   int       `json:"capacity"`
	Building   *Building `json:"building"`
	Students   *Students `json:"students"`
}
type Building struct {
	Street string `json:"street"`
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
