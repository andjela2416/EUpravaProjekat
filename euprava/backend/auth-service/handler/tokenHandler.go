package handlers

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

func ExchangeCodeForToken(code string) (*oauth2.Token, error) {
	oauth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: keycloakIssuer + "/protocol/openid-connect/token",
		},
		RedirectURL: redirectURL,
	}

	return oauth2Config.Exchange(context.Background(), code)
}

func ValidateJWT(tokenString string) (*UserInfo, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Claims.(jwt.MapClaims)["aud"].(string) != allowedAudience {
			return nil, fmt.Errorf("invalid token audience")
		}
		return []byte(clientSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
	}, nil
}
