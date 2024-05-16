package data

import (
	"encoding/json"
	"io"
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

type Application struct {
	Status  string   `json:"status"` //accepted / rejected / pending
	Student *Student `json:"student"`
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
