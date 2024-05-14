package main

import (
	"log"
	"os"

	h "auth-service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", h.ServeLoginPage)
	router.GET("/login", h.RedirectToKeycloak)
	router.GET("/callback", h.HandleKeycloakCallback)

	port := os.Getenv("PORT")
	if port == "" {
		port = "18080"
	}
	log.Printf("Server listening on port %s...", port)
	log.Fatal(router.Run(":" + port))
}
