package main

import (
	"log"
	"os"

	h "auth-service/handler"

	"github.com/gin-gonic/gin"
)

var (
	keycloakIssuer  = "https://your-keycloak-server/auth/realms/your-realm"
	clientID        = "your-client-id"
	clientSecret    = "your-client-secret"
	redirectURL     = "http://localhost:8080/callback"
	allowedAudience = "your-client-id"
)

func main() {
	router := gin.Default()

	router.GET("/", h.ServeLoginPage)
	router.GET("/login", h.RedirectToKeycloak)
	router.GET("/callback", h.HandleKeycloakCallback)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s...", port)
	log.Fatal(router.Run(":" + port))
}
