package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"time"
)

type Student struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Firstname      string             `bson:"firstName,omitempty" json:"firstName,omitempty"`
	Lastname       string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Gender         Gender             `bson:"gender,omitempty" json:"gender,omitempty"`
	DateOfBirth    int                `bson:"date_of_birth,omitempty" json:"date_of_birth,omitempty"`
	Residence      string             `bson:"residence,omitempty" json:"residence,omitempty"`
	Email          string             `bson:"email,omitempty" json:"email,omitempty"`
	Username       string             `bson:"username,omitempty" json:"username,omitempty"`
	UserType       UserType           `bson:"userType,omitempty" json:"userType,omitempty"`
	HealthRecordID primitive.ObjectID `bson:"healthRecordID,omitempty" json:"healthRecordID,omitempty"`
}

type Gender string

const (
	Male   = "Male"
	Female = "Female"
)

type Students []*Student

type TherapyData struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
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

// predstavlja pregled pacijenta
type AppointmentData struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	StudentID    primitive.ObjectID `bson:"student_id,omitempty" json:"student_id,omitempty"`
	Date         time.Time          `bson:"date,omitempty" json:"date,omitempty"`
	DoorNumber   int                `bson:"door_number" json:"door_number"`
	Description  string             `bson:"description" json:"description"`
	Systematic   bool               `bson:"systematic" json:"systematic"`
	FacultyName  string             `bson:"faculty_name" json:"faculty_name"`
	FieldOfStudy string             `bson:"field_of_study" json:"field_of_study"`
	Reserved     bool               `bson:"reserved" json:"reserved"`
}

type Appointments []*AppointmentData

type Medication struct {
	Name      string `bson:"name,omitempty" json:"name,omitempty"`
	Dosage    string `bson:"dosage,omitempty" json:"dosage,omitempty"`
	Frequency string `bson:"frequency,omitempty" json:"frequency,omitempty"`
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

func (o *AppointmentData) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *AppointmentData) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *Appointments) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Appointments) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

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

type HealthRecord struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	StudentID  primitive.ObjectID `bson:"studentId,omitempty" json:"studentId,omitempty"`
	RecordData string             `bson:"recordData,omitempty" json:"recordData,omitempty"`
}

type HealthRecords []*HealthRecord

func (o *HealthRecord) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *HealthRecord) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *HealthRecords) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *HealthRecords) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

type UserType string

const (
	STUDENT = "STUDENT"
	DOCTOR  = "DOCTOR"
)

type UsernameChange struct {
	OldUsername string `json:"old_username"`
	NewUsername string `json:"new_username"`
}
