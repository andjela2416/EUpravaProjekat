package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (r *HealthCareHandler) InsertUser(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyProduct{}).(*data.User)
	err := r.healthCareRepo.InsertUser(user)
	if err != nil {
		r.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error creating user."))
	}
	rw.WriteHeader(http.StatusOK)
}

func (h *HealthCareHandler) UpdateHealthRecord(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hRecordId := vars["id"]

	hRecord := r.Context().Value(KeyProduct{}).(*data.HealthRecord)

	err := h.healthCareRepo.UpdateHealthRecord(hRecordId, hRecord)
	if err != nil {
		h.logger.Println("Error updating health record:", err)
		http.Error(rw, "Error updating health record", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (r *HealthCareHandler) GetAllUsers(rw http.ResponseWriter, h *http.Request) {
	users, err := r.healthCareRepo.GetAllUsers()
	if err != nil {
		r.logger.Print("Database exception")
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		r.logger.Fatal("Unable to convert to json")
		return
	}
}

func (r *HealthCareHandler) GetAllHealthRecords(rw http.ResponseWriter, h *http.Request) {
	users, err := r.healthCareRepo.GetAllHealthRecords()
	if err != nil {
		r.logger.Print("Database exception")
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		r.logger.Fatal("Unable to convert to json")
		return
	}
}

// GetStudentByID vraća podatke o studentu po njegovom ID-ju.
func (h *HealthCareHandler) GetUserByID(rw http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(rw, "Missing User ID", http.StatusBadRequest)
		return
	}

	user, err := h.healthCareRepo.GetUserByID(userID)
	if err != nil {
		h.logger.Print("Database exception: ", err)
		http.Error(rw, "Error retrieving user.", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(rw, "user not found.", http.StatusNotFound)
		return
	}

	err = user.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error encoding user data.", http.StatusInternalServerError)
		return
	}
}

func (h *HealthCareHandler) GetHealthRecordByID(rw http.ResponseWriter, r *http.Request) {
	hRecordID := r.URL.Query().Get("id")
	if hRecordID == "" {
		http.Error(rw, "Missing health record ID", http.StatusBadRequest)
		return
	}

	hRecord, err := h.healthCareRepo.GetHealthRecordByID(hRecordID)
	if err != nil {
		h.logger.Print("Database exception: ", err)
		http.Error(rw, "Error retrieving hRecord.", http.StatusInternalServerError)
		return
	}

	if hRecord == nil {
		http.Error(rw, "hRecord not found.", http.StatusNotFound)
		return
	}

	err = hRecord.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error encoding hRecord data.", http.StatusInternalServerError)
		return
	}
}

// UpdateStudent ažurira podatke o studentu.
func (r *HealthCareHandler) UpdateUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	user := h.Context().Value(KeyProduct{}).(*data.User)

	err := r.healthCareRepo.UpdateUser(id, user)
	if err != nil {
		http.Error(rw, "Error updating user", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// DeleteStudent briše studenta iz baze podataka.
func (h *HealthCareHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(rw, "Missing user ID", http.StatusBadRequest)
		return
	}

	err := h.healthCareRepo.DeleteUser(userID)
	if err != nil {
		h.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error deleting user."))
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (h *HealthCareHandler) CreateAppointment(rw http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("doctorId")
	if userID == "" {
		http.Error(rw, "User ID is required", http.StatusBadRequest)
		return
	}

	// Konverzija userID u ObjectID
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		h.logger.Println("Invalid user ID:", err)
		http.Error(rw, "Invalid user ID", http.StatusBadRequest)
		return
	}

	appointmentData := r.Context().Value(KeyProduct{}).(*data.AppointmentData)
	fmt.Printf("Received appointment: %+v\n", appointmentData)

	appointmentData.DoctorID = oid

	err = h.healthCareRepo.CreateAppointment(r, appointmentData)
	if err != nil {
		h.logger.Print("Database exception: ", err)
		http.Error(rw, "Error creating appointment.", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

// GetAppointmentByID vraća pregled po ID-u.
func (h *HealthCareHandler) GetAppointmentByID(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	appointmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(rw, "Invalid ID format", http.StatusBadRequest)
		return
	}

	appointment, err := h.healthCareRepo.GetAppointmentByID(appointmentID)
	if err != nil {
		http.Error(rw, "Error retrieving appointment", http.StatusInternalServerError)
		return
	}

	if appointment == nil {
		http.Error(rw, "Appointment not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(rw).Encode(appointment)
	if err != nil {
		http.Error(rw, "Error encoding appointment to JSON", http.StatusInternalServerError)
		return
	}
}

func (r *HealthCareHandler) UpdateAppointment(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	appointment := h.Context().Value(KeyProduct{}).(*data.AppointmentData)

	err := r.healthCareRepo.UpdateAppointment(id, appointment)
	if err != nil {
		http.Error(rw, "Error updating appointment", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// DeleteAppointment briše pregled po ID-u.
func (h *HealthCareHandler) DeleteAppointment(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := h.healthCareRepo.DeleteAppointment(id)
	if err != nil {
		http.Error(rw, "Error deleting appointment", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h *HealthCareHandler) ScheduleAppointment(rw http.ResponseWriter, r *http.Request) {
	var request struct {
		AppointmentID string `json:"appointment_id"`
		UserID        string `json:"user_id"` // Dodajte UserID
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Println("Error decoding request:", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid request payload"))
		return
	}

	appointmentID, err := primitive.ObjectIDFromHex(request.AppointmentID)
	if err != nil {
		h.logger.Println("Invalid appointment ID:", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid appointment ID"))
		return
	}

	userID, err := primitive.ObjectIDFromHex(request.UserID)
	if err != nil {
		h.logger.Println("Invalid user ID:", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid user ID"))
		return
	}

	err = h.healthCareRepo.ScheduleAppointment(r, appointmentID, userID)
	if err != nil {
		h.logger.Println("Error scheduling appointment:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error scheduling appointment"))
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// GetAllReservedAppointmentsForUser vraća sve rezervisane termine pregleda za određenog korisnika.
func (h *HealthCareHandler) GetAllReservedAppointmentsForUser(rw http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(rw, "User ID is required", http.StatusBadRequest)
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		h.logger.Print("Invalid user ID format: ", err)
		http.Error(rw, "Invalid user ID", http.StatusBadRequest)
		return
	}

	appointments, err := h.healthCareRepo.GetAllReservedAppointmentsForUser(userID)
	if err != nil {
		h.logger.Print("Database exception: ", err)
		http.Error(rw, "Error retrieving reserved appointments.", http.StatusInternalServerError)
		return
	}

	if appointments == nil {
		http.Error(rw, "No appointments found for the user.", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(rw).Encode(appointments)
	if err != nil {
		h.logger.Print("Error encoding appointments to JSON: ", err)
		http.Error(rw, "Error encoding appointments to JSON.", http.StatusInternalServerError)
		return
	}
}

func (h *HealthCareHandler) GetAllAppointmentsForUser(rw http.ResponseWriter, r *http.Request) {

	appointments, err := h.healthCareRepo.GetAllAppointmentsForUser(r)
	if err != nil {
		h.logger.Print("Database exception: ", err)
		http.Error(rw, "Error retrieving appointments.", http.StatusInternalServerError)
		return
	}

	if appointments == nil {
		http.Error(rw, "No appointments found for the user.", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(rw).Encode(appointments)
	if err != nil {
		h.logger.Print("Error encoding appointments to JSON: ", err)
		http.Error(rw, "Error encoding appointments to JSON.", http.StatusInternalServerError)
		return
	}
}

func (h *HealthCareHandler) CancelAppointment(rw http.ResponseWriter, r *http.Request) {
	var request struct {
		AppointmentID string `json:"appointment_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Println("Error decoding request:", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid request payload"))
		return
	}

	appointmentID, err := primitive.ObjectIDFromHex(request.AppointmentID)
	if err != nil {
		h.logger.Println("Invalid appointment ID:", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid appointment ID"))
		return
	}

	err = h.healthCareRepo.CancelAppointment(appointmentID)
	if err != nil {
		h.logger.Println("Error canceling appointment:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error canceling appointment"))
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// GetAllReservedAppointments vraća sve rezervisane termine pregleda.
func (h *HealthCareHandler) GetAllReservedAppointments(rw http.ResponseWriter, r *http.Request) {
	appointments, err := h.healthCareRepo.GetAllReservedAppointments()
	if err != nil {
		http.Error(rw, "Error retrieving appointments", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(appointments)
	if err != nil {
		http.Error(rw, "Error encoding appointments to JSON", http.StatusInternalServerError)
		return
	}
}

// GetAllNotReservedAppointments vraća sve nerezevisane termine pregleda.
func (h *HealthCareHandler) GetAllNotReservedAppointments(rw http.ResponseWriter, r *http.Request) {
	appointments, err := h.healthCareRepo.GetAllNotReservedAppointments()
	if err != nil {
		http.Error(rw, "Error retrieving appointments", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(appointments)
	if err != nil {
		http.Error(rw, "Error encoding appointments to JSON", http.StatusInternalServerError)
		return
	}
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
	_, err := h.healthCareRepo.SaveTherapyData(therapyData)
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

	_, err = h.healthCareRepo.SaveTherapyData(&therapyData)
	if err != nil {
		h.logger.Println("Error saving therapy data:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

// UpdateTherapyData ažurira podatke o terapiji u bazi podataka.
func (r *HealthCareHandler) UpdateTherapyData(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	therapy := h.Context().Value(KeyProduct{}).(*data.TherapyData)

	err := r.healthCareRepo.UpdateTherapyData(id, therapy)
	if err != nil {
		http.Error(rw, "Error updating appointment", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// DeleteTherapyData briše podatke o terapiji iz baze podataka.
func (h *HealthCareHandler) DeleteTherapyData(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	therapyIDHex := vars["id"]

	therapyID, err := primitive.ObjectIDFromHex(therapyIDHex)
	if err != nil {
		http.Error(rw, "Invalid therapy ID", http.StatusBadRequest)
		return
	}

	err = h.healthCareRepo.DeleteTherapyData(therapyID)
	if err != nil {
		http.Error(rw, "Error deleting therapy data", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// GetTherapyDataByID vraća podatke o terapiji na osnovu ID-a.
func (h *HealthCareHandler) GetTherapyDataByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	therapyIDHex := vars["id"]

	therapyID, err := primitive.ObjectIDFromHex(therapyIDHex)
	if err != nil {
		http.Error(rw, "Invalid therapy ID", http.StatusBadRequest)
		return
	}

	therapyData, err := h.healthCareRepo.GetTherapyDataByID(therapyID)
	if err != nil {
		http.Error(rw, "Error getting therapy data", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(therapyData)
	if err != nil {
		http.Error(rw, "Error encoding therapy data to JSON", http.StatusInternalServerError)
		return
	}
}

func (h *HealthCareHandler) SaveAndShareTherapyDataWithDietService(rw http.ResponseWriter, r *http.Request) {
	therapyData := r.Context().Value(KeyProduct{}).(*data.TherapyData)

	therapyData.Status = "SentToFoodService"

	err := h.healthCareRepo.SaveAndShareTherapyDataWithDietService(therapyData)
	if err != nil {
		h.logger.Print("Error sharing therapy data with diet service: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error sharing therapy data with diet service."))
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h *HealthCareHandler) UpdateTherapyFromFoodService(rw http.ResponseWriter, r *http.Request) {
	therapyData := r.Context().Value(KeyProduct{}).(*data.TherapyData)

	err := h.healthCareRepo.UpdateTherapyFromFoodService(therapyData)
	if err != nil {
		h.logger.Println("Error updating therapy from food service:", err)
		rw.WriteHeader(http.StatusInternalServerError)
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

func (s *HealthCareHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		users := &data.User{}
		err := users.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			s.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, users)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}

func (s *HealthCareHandler) MiddlewareAuthUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		users := &data.AuthUser{}
		err := users.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			s.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, users)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}

func (s *HealthCareHandler) MiddlewareAppointmentDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		appointments := &data.AppointmentData{}
		err := appointments.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			s.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, appointments)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}

func (s *HealthCareHandler) MiddlewareTherapyDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		therapies := &data.TherapyData{}
		err := therapies.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			s.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, therapies)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}

func (s *HealthCareHandler) MiddlewareHealthRecordDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		healthRecords := &data.HealthRecord{}
		err := healthRecords.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			s.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, healthRecords)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}
