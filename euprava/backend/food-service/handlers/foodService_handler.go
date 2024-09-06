package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"food-service/data"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FoodServiceHandler struct {
	logger          *log.Logger
	foodServiceRepo *data.FoodServiceRepo
}

type KeyProduct struct{}
type KeyFood struct{}

func NewFoodServiceHandler(l *log.Logger, r *data.FoodServiceRepo) *FoodServiceHandler {
	return &FoodServiceHandler{l, r}
}

// GetListFoodHandler vraća sve unose hrane iz baze
func (h *FoodServiceHandler) GetListFoodHandler(rw http.ResponseWriter, r *http.Request) {
	foodList, err := h.foodServiceRepo.GetListFood()
	if err != nil {
		h.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error retrieving food entries."))
		return
	}

	// Konvertuj listu hrane u JSON i pošalji klijentu
	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(foodList)
	if err != nil {
		h.logger.Print("Error converting food list to JSON: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error converting food list to JSON."))
		return
	}
}

// CreateFoodHandler kreira novi unos hrane sa stanjem postavljenim na 'Neporucena'
func (h *FoodServiceHandler) CreateFoodHandler(rw http.ResponseWriter, r *http.Request) {
	// Preuzmi podatke o hrani iz konteksta
	foodData := r.Context().Value(KeyFood{}).(*data.Food)

	fmt.Printf("Received food entry: %+v\n", foodData)

	// Kreiraj novi unos hrane koristeći metodu iz repo
	err := h.foodServiceRepo.CreateFoodEntry(r, foodData)
	if err != nil {
		h.logger.Print("Database exception: ", err)
		http.Error(rw, "Error creating food entry.", http.StatusInternalServerError)
		return
	}

	// Postavi status kod 201 Created
	rw.WriteHeader(http.StatusCreated)

	// Definiši strukturu odgovora sa porukom i podacima o kreiranoj hrani
	response := map[string]interface{}{
		"message":  "Food entry created successfully",
		"foodName": foodData.FoodName,
		"status":   foodData.Stanje,
	}

	// Pošalji JSON odgovor nazad klijentu
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
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

func (h *FoodServiceHandler) ClearTherapiesList(rw http.ResponseWriter, r *http.Request) {
	// Poziv funkcije za brisanje liste terapija iz repozitorijuma
	err := h.foodServiceRepo.ClearTherapiesCache()
	if err != nil {
		h.logger.Print("Error clearing therapies list: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error clearing therapies list."))
		return
	}
	// Odgovor sa statusom OK ako je brisanje uspeÄąË‡no
	rw.WriteHeader(http.StatusOK)
}

func (h *FoodServiceHandler) UpdateTherapyStatus(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	therapyID := vars["id"]
	//status := r.FormValue("status")

	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(rw, "Invalid request body: unable to decode JSON", http.StatusBadRequest)
	}

	status, ok := requestBody["status"].(string)
	if !ok {
		log.Println("Status nije string")
	}

	fmt.Println("Received status:", status)

	switch status {
	case data.Done, data.Undone:
		objectID, err := primitive.ObjectIDFromHex(therapyID)
		if err != nil {
			h.logger.Printf("Invalid therapy ID provided: %v", err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Invalid therapy ID provided."))
			return
		}

		// Call repository to update therapy status in cache
		err = h.foodServiceRepo.UpdateTherapyStatus(objectID, data.Status(status))
		if err != nil {
			h.logger.Printf("Error updating therapy status: %v", err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Error updating therapy status."))
			return
		}
		// Respond with success status
		rw.WriteHeader(http.StatusOK)
	default:
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid status provided."))
		return
	}
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

// MiddlewareFoodDeserialization je middleware funkcija koja preuzima podatke o hrani iz zahteva i stavlja ih u kontekst
func (h *FoodServiceHandler) MiddlewareFoodDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		foodData := &data.Food{}

		// Pokušaj deserializacije JSON podataka iz tela zahteva
		err := json.NewDecoder(r.Body).Decode(foodData)
		if err != nil {
			http.Error(rw, "Unable to decode JSON", http.StatusBadRequest)
			h.logger.Fatal(err)
			return
		}

		// Postavi deserializovane podatke u kontekst
		ctx := context.WithValue(r.Context(), KeyFood{}, foodData)
		r = r.WithContext(ctx)

		// Nastavi sa sledeæim handlerom
		next.ServeHTTP(rw, r)
	})
}
