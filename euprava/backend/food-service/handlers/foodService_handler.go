package handlers

import (
	"context"
	"food-service/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type FoodServiceHandler struct {
	logger          *log.Logger
	foodServiceRepo *data.FoodServiceRepo
}

type KeyProduct struct{}

func NewFoodServiceHandler(l *log.Logger, r *data.FoodServiceRepo) *FoodServiceHandler {
	return &FoodServiceHandler{l, r}
}

// mongo
// editFood
func (r *FoodServiceHandler) EditFoodForStudent(rw http.ResponseWriter, h *http.Request) {
	// Parse request parameters
	vars := mux.Vars(h)
	studentID := vars["id"]        // Presuming "id" is the parameter name for student ID
	newFood := h.FormValue("food") // Assuming "food" is the parameter name for new food

	// Call repository to edit food for student
	err := r.foodServiceRepo.EditFoodForStudent(studentID, newFood)
	if err != nil {
		r.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error updating student's food."))
		return
	}

	// Respond with success status
	rw.WriteHeader(http.StatusOK)
}

// getAllFood
func (r *FoodServiceHandler) GetAllFoodOfStudents(rw http.ResponseWriter, h *http.Request) {
	students, err := r.foodServiceRepo.GetAllFoodOfStudents()
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

func (r *FoodServiceHandler) GetTherapiesFromHealthCare(rw http.ResponseWriter, h *http.Request) {
	therapies, err := r.foodServiceRepo.GetAllTherapiesFromHealthCareService()
	if err != nil {
		r.logger.Print("Database exception")
	}

	if therapies == nil {
		return
	}

	err = therapies.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		r.logger.Fatal("Unable to convert to json")
		return
	}
}

func (r *FoodServiceHandler) GetTherapies(rw http.ResponseWriter, h *http.Request) {
	therapies, err := r.foodServiceRepo.GetAllTherapiesFromFoodService()
	if err != nil {
		r.logger.Print("Error retrieving therapies from Food Service:", err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	if therapies == nil {
		http.Error(rw, "No therapies found", http.StatusNotFound)
		return
	}

	err = therapies.ToJSON(rw)
	if err != nil {
		r.logger.Print("Error converting therapies to JSON:", err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *FoodServiceHandler) SaveTherapy(rw http.ResponseWriter, r *http.Request) {
	therapyData := r.Context().Value(KeyProduct{}).(*data.TherapyData)
	err := h.foodServiceRepo.SaveTherapyData(therapyData)
	if err != nil {
		h.logger.Print("Error sharing therapy data with diet service: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error sharing therapy data with diet service."))
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (s *FoodServiceHandler) MiddlewareStudentDeserialization(next http.Handler) http.Handler {
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

func (s *FoodServiceHandler) MiddlewareTherapyDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		students := &data.TherapyData{}
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
