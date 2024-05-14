package handlers

import (
	"context"
	"encoding/json"
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

func (h *HealthCareHandler) ScheduleAppointment(rw http.ResponseWriter, r *http.Request) {
	appointmentData := r.Context().Value(KeyProduct{}).(*data.AppointmentData)
	err := h.healthCareRepo.ScheduleAppointment(appointmentData)
	if err != nil {
		h.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error scheduling appointment."))
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (h *HealthCareHandler) GetAllAppointments(rw http.ResponseWriter, r *http.Request) {
	appointments, err := h.healthCareRepo.GetAllAppointments()
	if err != nil {
		h.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error retrieving appointments."))
		return
	}

	err = appointments.ToJSON(rw)
	if err != nil {
		h.logger.Print("Error converting appointments to JSON: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error converting appointments to JSON."))
		return
	}
}

func (h *HealthCareHandler) GetAllTherapies(rw http.ResponseWriter, r *http.Request) {
	therapies, err := h.healthCareRepo.GetAllTherapies()
	if err != nil {
		h.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error retrieving therapies."))
		return
	}

	err = therapies.ToJSON(rw)
	if err != nil {
		h.logger.Print("Error converting therapies to JSON: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error converting therapies to JSON."))
		return
	}
}

func (h *HealthCareHandler) SaveTherapy(rw http.ResponseWriter, r *http.Request) {
	therapyData := r.Context().Value(KeyProduct{}).(*data.TherapyData)
	err := h.healthCareRepo.SaveTherapyData(therapyData)
	if err != nil {
		h.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error writing therapy from examination."))
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (h *HealthCareHandler) SaveTherapyData(rw http.ResponseWriter, r *http.Request) {
	var therapyData data.TherapyData
	err := json.NewDecoder(r.Body).Decode(&therapyData)
	if err != nil {
		h.logger.Println("Error decoding therapy data:", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.healthCareRepo.SaveTherapyData(&therapyData)
	if err != nil {
		h.logger.Println("Error saving therapy data:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (h *HealthCareHandler) SaveAndShareTherapyDataWithDietService(rw http.ResponseWriter, r *http.Request) {
	therapyData := r.Context().Value(KeyProduct{}).(*data.TherapyData)
	err := h.healthCareRepo.SaveAndShareTherapyDataWithDietService(therapyData)
	if err != nil {
		h.logger.Print("Error sharing therapy data with diet service: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error sharing therapy data with diet service."))
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (h *HealthCareHandler) GetDoneTherapiesFromFoodService(rw http.ResponseWriter, r *http.Request) {
	doneTherapies, err := h.healthCareRepo.GetDoneTherapiesFromFoodService()
	if err != nil {
		h.logger.Printf("Error getting done therapies from food service: %v\n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error getting done therapies from food service"))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(doneTherapies)
	if err != nil {
		h.logger.Printf("Error encoding done therapies to JSON: %v\n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error encoding done therapies to JSON"))
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

func (s *HealthCareHandler) MiddlewareAppointmentDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		students := &data.AppointmentData{}
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

func (s *HealthCareHandler) MiddlewareTherapyDeserialization(next http.Handler) http.Handler {
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
