package repositories

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserType string
type Gender string

const (
	StudentType        UserType = "STUDENT"
	ProfessorType      UserType = "PROFESSOR"
	AdministratorType  UserType = "ADMINISTRATOR"
	StudentServiceType UserType = "STUDENTSKA_SLUZBA"

	Male   Gender = "Male"
	Female Gender = "Female"
)

type University struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name" json:"name" validate:"required"`
	Location        string             `bson:"location" json:"location"`
	FoundationYear  int                `bson:"foundation_year" json:"foundation_year"`
	StudentCount    int                `bson:"student_count" json:"student_count"`
	StaffCount      int                `bson:"staff_count" json:"staff_count"`
	Accreditation   string             `bson:"accreditation" json:"accreditation"`
	OfferedPrograms []string           `bson:"offered_programs" json:"offered_programs"`
}

type Department struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name" validate:"required"`
	Chief         Professor          `bson:"chief" json:"chief"`
	StudyPrograms []string           `bson:"study_programs" json:"study_programs"`
	Staff         []Professor        `bson:"staff" json:"staff"`
}

type Professor struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Subjects []Course           `bson:"subjects" json:"subjects"`
	Office   string             `bson:"office" json:"office"`
}

type Assistant struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Professor Professor          `bson:"professor" json:"professor"`
	Courses   []Course           `bson:"courses" json:"courses"`
}

type Course struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name" validate:"required"`
	Department    Department         `bson:"department" json:"department"`
	Professor     Professor          `bson:"professor" json:"professor"`
	Schedule      string             `bson:"schedule" json:"schedule"`
	Prerequisites []string           `bson:"prerequisites" json:"prerequisites"`
}

type StudentService struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name" validate:"required"`
	Department    string             `bson:"department" json:"department"`
	EmployeeCount int                `bson:"employee_count" json:"employee_count"`
}

type Administrator struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Position       string             `bson:"position" json:"position"`
	StudentService StudentService     `bson:"student_service" json:"student_service"`
}

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName      string             `bson:"first_name" json:"first_name" validate:"required"`
	LastName       string             `bson:"last_name" json:"last_name"`
	Username       string             `bson:"username" json:"username"`
	Email          string             `bson:"email" json:"email" validate:"required,email"`
	DateOfBirth    time.Time          `bson:"date_of_birth" json:"date_of_birth"`
	Password       string             `bson:"password" json:"password"`
	UserType       UserType           `bson:"user_type" json:"user_type"`
	StudentDetails *Student           `bson:"student_details,omitempty" json:"student_details,omitempty"`
}

type Student struct {
	User
	Major         string  `bson:"major" json:"major,omitempty"`
	Year          int     `bson:"year" json:"year,omitempty"`
	AssignedDorm  string  `bson:"assigned_dorm" json:"assigned_dorm,omitempty"`
	Scholarship   bool    `bson:"scholarship" json:"scholarship,omitempty"`
	HighschoolGPA float64 `bson:"highschool_gpa" json:"highschool_gpa,omitempty"`
	GPA           float64 `bson:"gpa" json:"gpa,omitempty"`
	ESBP          int     `bson:"esbp" json:"esbp,omitempty"`
}

type Exam struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Student  Student            `bson:"student" json:"student"`
	Course   Course             `bson:"course" json:"course"`
	ExamDate time.Time          `bson:"exam_date" json:"exam_date"`
	Status   string             `bson:"status" json:"status"`
}

type TuitionPayment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StudentID primitive.ObjectID `bson:"student_id" json:"student_id"`
	Amount    float64            `bson:"amount" json:"amount"`
	Date      time.Time          `bson:"date" json:"date"`
}

type Notification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title" validate:"required"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

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

func (u *University) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (d *Department) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (p *Professor) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (a *Assistant) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (c *Course) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (ss *StudentService) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ss)
}

func (adm *Administrator) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(adm)
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (s *Student) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

func (e *Exam) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(e)
}

func (u *University) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

func (d *Department) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(d)
}

func (p *Professor) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (a *Assistant) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

func (c *Course) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c)
}

func (ss *StudentService) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(ss)
}

func (adm *Administrator) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(adm)
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

func (s *Student) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(s)
}

func (e *Exam) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(e)
}
