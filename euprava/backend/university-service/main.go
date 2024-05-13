package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8001"
	}

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
