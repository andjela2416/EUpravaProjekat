package models

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
