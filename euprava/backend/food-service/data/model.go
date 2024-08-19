package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Student struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Firstname string             `bson:"firstName,omitempty" json:"firstName,omitempty"`
	Lastname  string             `bson:"lastName,omitempty" json:"lastName,omitempty"`

	//UserType  UserType           `bson:"userType" json:"userType"`
}

type Students []*Student

type TherapyData struct {
	ID        primitive.ObjectID `bson:"therapyId,omitempty" json:"therapyId,omitempty"`
	StudentID primitive.ObjectID `bson:"studentId,omitempty" json:"studentId,omitempty"`
	Diagnosis string             `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Status    Status             `bson:"status,omitempty" json:"status,omitempty"`
	//Medications  []Medication       `bson:"medications,omitempty" json:"medications,omitempty"`
	//Instructions string             `bson:"instructions,omitempty" json:"instructions,omitempty"`
}

type Status string

const (
	SentToFoodService = "sent to food service"
	Done              = "done"
	Undone            = "undone"
)

type Therapies []*TherapyData

func (o *TherapyData) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *TherapyData) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *Therapies) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Therapies) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *Students) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Students) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *Student) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Student) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

type UserType string

const (
	Guest = "Guest"
	Host  = "Host"
)

type UsernameChange struct {
	OldUsername string `json:"old_username"`
	NewUsername string `json:"new_username"`
}
