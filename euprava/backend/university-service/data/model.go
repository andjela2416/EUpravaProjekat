package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	StudentId string             `bson:"student_id" json:"student_id"` //SRXX/20XX : SR22/2022
	Firstname string             `bson:"firstName,omitempty" json:"firstName,omitempty"`
	Lastname  string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Gender    Gender             `bson:"gender,omitempty" json:"gender,omitempty"`
	Age       int                `bson:"age,omitempty" json:"age,omitempty"`
	Residence string             `bson:"residence,omitempty" json:"residence,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Username  string             `bson:"username" json:"username"`
	//UserType  UserType           `bson:"userType" json:"userType"`
	StudyInfo StudyInfo `bson:"study_info,omitempty" json:"study_info,omitempty"`
}

type StudyInfo struct {
	HighschoolGPA float64 `bson:"highschool_gpa,omitempty" json:"highschool_gpa,omitempty"`
	GPA           float64 `bson:"gpa,omitempty" json:"gpa,omitempty"`
	ESBP          int     `bson:"esbp,omitempty" json:"esbp,omitempty"`
	Year          int     `bson:"year,omitempty" json:"year,omitempty"`
}

type Gender string

const (
	Male   = "Male"
	Female = "Female"
)

type Students []*Student
