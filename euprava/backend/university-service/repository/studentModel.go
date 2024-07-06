package repositories

import (
	"encoding/json"
	"time"

	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	First_Name *string            `bson:"first_name" json:"first_name" validate:"required"`
	Last_Name  *string            `bson:"last_name" json:"last_name"`
	Gender     Gender             `bson:"gender,omitempty" json:"gender,omitempty"`
	Residence  string             `bson:"residence,omitempty" json:"residence,omitempty"`
	Username   *string            `bson:"username" json:"username"`
	Email      string             `bson:"email" json:"email" validate:"required,email"`
	Address    *string            `bson:"address" json:"address"`
	User_type  *string            `bson:"user_type" json:"user_type" validate:"required,oneof=Guest Host User"`
	StudyInfo  StudyInfo          `bson:"study_info,omitempty" json:"study_info,omitempty"`
}

type StudyInfo struct {
	HighschoolGPA float64 `bson:"highschool_gpa,omitempty" json:"highschool_gpa,omitempty"`
	GPA           float64 `bson:"gpa,omitempty" json:"gpa,omitempty"`
	ESBP          int     `bson:"esbp,omitempty" json:"esbp,omitempty"`
	Year          int     `bson:"year,omitempty" json:"year,omitempty"`
}

type Notification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title" validate:"required"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type Gender string

const (
	Male   = "Male"
	Female = "Female"
)

type Students []*Student

type Notifications []*Notification

func (n *Notifications) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(n)
}

func (n *Notification) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(n)
}

func (n *Notification) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(n)
}

func (s *Students) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

func (s *Student) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

func (s *Student) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(s)
}
