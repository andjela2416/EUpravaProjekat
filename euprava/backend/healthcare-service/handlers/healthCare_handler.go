package handlers

import (
	"context"
	"healthcare-service/data"
	"log"
	"net/http"
)

type HealthCareHandler struct {
	logger         *log.Logger
	healthCareRepo *data.HealthCareRepo
}

type KeyProduct struct{}

func NewHealthCareHandler(l *log.Logger, r *data.HealthCareRepo) *HealthCareHandler {
	return &HealthCareHandler{l, r}
}

// mongo
func (r *HealthCareHandler) InsertStudent(rw http.ResponseWriter, h *http.Request) {
	student := h.Context().Value(KeyProduct{}).(*data.Student)
	err := r.healthCareRepo.InsertStudent(student)
	if err != nil {
		r.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error creating student."))
	}
	rw.WriteHeader(http.StatusOK)
}

func (r *HealthCareHandler) GetAllStudents(rw http.ResponseWriter, h *http.Request) {
	students, err := r.healthCareRepo.GetAllStudents()
	if err != nil {
		r.logger.Print("Database exception")
	}

	if students == nil {
		return
	}

	err = students.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		r.logger.Fatal("Unable to convert to json")
		return
	}
}

func (s *HealthCareHandler) MiddlewareStudentDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		students := &data.Student{}
		err := students.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			s.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, students)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}
