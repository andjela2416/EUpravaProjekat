package main

import (
	"os"
	"time"
	"university-service/routes"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8088"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// CORS
	router.Use(cors.Middleware(cors.Config{
		Origins:         "http://localhost:3000, *",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	routes.ProfileRoutes(router)

	router.Run(":" + port)
}
