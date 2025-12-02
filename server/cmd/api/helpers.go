package main

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/corbinlazarone/Todue-Actual/cmd/internals/models"
	"github.com/golang-jwt/jwt/v5"
)

func extractBearerToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is empty")
	}

	// Check for "Bearer " prefix
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

func generateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":        user.Id,
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(1 * time.Hour).Unix(), // Expires 1 hour
		"iss":        "todue-api",
		"aud":        "todue-web",
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"picture":    user.Picture,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Fatal("JWT_SECRET not set")
	}

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
