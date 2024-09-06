package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"time"
)

type User struct {
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

type Users []*User

type TherapyData struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	StudentHealthRecordID primitive.ObjectID `bson:"studentHealthRecordID,omitempty" json:"studentHealthRecordID,omitempty"`
	Diagnosis             string             `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Status                Status             `bson:"status,omitempty" json:"status,omitempty"`
	//Medications  []Medication       `bson:"medications,omitempty" json:"medications,omitempty"`
	//Instructions string             `bson:"instructions,omitempty" json:"instructions,omitempty"`
}

type AuthUser struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Email         *string            `json:"email" validate:"email,required"`
	Password      *string            `json:"password" validate:"required,min=8"`
	Phone         *string            `json:"phone" validate:"required"`
	Address       *string            `json:"address" validate:"required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate:"required"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
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
	DoctorID     primitive.ObjectID `bson:"doctor_id" json:"doctor_id"`
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

func (o *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Users) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *AuthUser) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}

func (o *AuthUser) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *User) FromJSON(r io.Reader) error {
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
	UserID     primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`
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
