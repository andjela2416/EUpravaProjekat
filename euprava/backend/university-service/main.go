package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"

	"backend/university-service/database"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8001"
	}

	db, err := database.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database")
	}
	defer db.Close()
	fmt.Println("Connected to the database")

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
