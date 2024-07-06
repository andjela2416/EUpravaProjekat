package main

import (
	"context"
	"fmt"
	"food-service/data"
	"food-service/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World!")

	port := os.Getenv("FOOD_SERVICE_PORT")

	timeoutContext, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// Inicijalizacija loggera koji Ä‡e se koristiti, sa prefiksom i datumom za svaki log
	logger := log.New(os.Stdout, "[res-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[res-store] ", log.LstdFlags)

	// NoSQL: Inicijalizacija prodavnice proizvoda
	store, err := data.NewFoodServiceRepo(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.DisconnectMongo(timeoutContext)
	store.Ping()

	foodServiceHandler := handlers.NewFoodServiceHandler(logger, store)

	// Inicijalizacija rutera i dodavanje middleware-a za sve zahteve
	router := mux.NewRouter()
	router.Use(MiddlewareContentTypeSet)

	getAllFoodForStudents := router.Methods(http.MethodGet).Subrouter()
	getAllFoodForStudents.HandleFunc("/studentsfood", foodServiceHandler.GetAllFoodOfStudents)

	editFoodForStudent := router.Methods(http.MethodPost).Subrouter()
	editFoodForStudent.HandleFunc("/studentsfood", foodServiceHandler.EditFoodForStudent)
	editFoodForStudent.Use(foodServiceHandler.MiddlewareStudentDeserialization)

	getTherapies := router.Methods(http.MethodGet).Subrouter()
	getTherapies.HandleFunc("/therapies", foodServiceHandler.GetTherapies)

	saveTherapy := router.Methods(http.MethodPost).Subrouter()
	saveTherapy.HandleFunc("/therapy", foodServiceHandler.SaveTherapy)
	saveTherapy.Use(foodServiceHandler.MiddlewareTherapyDeserialization)

	clearAllTherapy := router.Methods(http.MethodDelete).Subrouter()
	clearAllTherapy.HandleFunc("/therapy", foodServiceHandler.ClearTherapiesList)

	updateTherapyStatus := router.Methods(http.MethodPut).Subrouter()
	updateTherapyStatus.HandleFunc("/therapy/{id}", foodServiceHandler.UpdateTherapyStatus)

	// Inicijalizacija HTTP servera
	server := http.Server{
		Addr:         ":" + port,
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Println("Server listening on port", port)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	if err := server.Shutdown(timeoutContext); err != nil {
		logger.Fatal("Cannot gracefully shutdown...", err)
	}
	logger.Println("Server stopped")
}

func MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		//s.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Set("X-Content-Type-Options", "nosniff")
		rw.Header().Set("X-Frame-Options", "DENY")
		rw.Header().Set("Content-Security-Policy", "script-src 'self' https://code.jquery.com https://cdn.jsdelivr.net https://www.google.com https://www.gstatic.com 'unsafe-inline' 'unsafe-eval'; style-src 'self' https://code.jquery.com https://cdn.jsdelivr.net https://fonts.googleapis.com https://fonts.gstatic.com 'unsafe-inline'; font-src 'self' https://code.jquery.com https://cdn.jsdelivr.net https://fonts.googleapis.com https://fonts.gstatic.com; img-src 'self' data: https://code.jquery.com https://i.ibb.co;")

		next.ServeHTTP(rw, h)
	})
}
