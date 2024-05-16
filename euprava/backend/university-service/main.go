package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"

	"university-service/data"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8001"
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[uni-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[uni-store] ", log.LstdFlags)

	store, err := data.NewUniRepo(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.DisconnectMongo(timeoutContext)
	store.Ping()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Middleware(cors.Config{
		Origins:         "https://localhost:3000, *",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Content-Type,Authorization",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.Run(":" + port)
}
