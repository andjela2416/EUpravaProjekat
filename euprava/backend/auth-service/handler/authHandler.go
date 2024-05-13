package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	keycloakIssuer  = "https://your-keycloak-server/auth/realms/your-realm"
	clientID        = "your-client-id"
	clientSecret    = "your-client-secret"
	redirectURL     = "http://localhost:8080/callback"
	allowedAudience = "your-client-id"
)

func ServeLoginPage(c *gin.Context) {
	c.Redirect(http.StatusFound, "http://localhost:8080/login")
}

func RedirectToKeycloak(c *gin.Context) {
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/protocol/openid-connect/auth?client_id=%s&redirect_uri=%s&response_type=code", keycloakIssuer, clientID, redirectURL))
}

func HandleKeycloakCallback(c *gin.Context) {
	code := c.Query("code")

	token, err := ExchangeCodeForToken(code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to exchange token")
		return
	}

	userInfo, err := ValidateJWT(token.AccessToken)
	if err != nil {
		c.String(http.StatusUnauthorized, "Failed to validate JWT token")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userInfo": userInfo,
	})
}

type UserInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
