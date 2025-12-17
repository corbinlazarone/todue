package main

import (
	"errors"
	"fmt"
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

func sanitizePDFText(input string) error {
	if len(input) > 100*1024 {
		return errors.New("input text too long")
	}

	// Forbidden keywords that could indicate prompt injection
	forbidden := []string{
		"system:",
		"user:",
		"assistant:",
		"ignore previous",
		"forget instructions",
		"override",
		"now do this",
		"respond with",
		"instead of",
		"bypass",
		"jailbreak",
	}

	lowerInput := strings.ToLower(input)
	for _, word := range forbidden {
		if strings.Contains(lowerInput, word) {
			return fmt.Errorf("input contains forbidden content: %s", word)
		}
	}

	// Character validation: only printable ASCII + basic Unicode
	for _, r := range input {
		if r < 32 && r != 9 && r != 10 && r != 13 { // allow tab, newline, carriage return
			return errors.New("input contains invalid characters")
		}
	}

	return nil
}

func ToRFC3339(dateStr, timeStr, tzStr string) (string, error) {
	loc, err := time.LoadLocation(tzStr)
	if err != nil {
		return "", err
	}

	datetime := dateStr + " " + timeStr
	t, err := time.ParseInLocation("2006-01-02 15:04", datetime, loc)
	if err != nil {
		return "", err
	}

	return t.Format(time.RFC3339), nil
}

// ValidateTimeZone validates that the timezone is a valid IANA timezone
func ValidateTimeZone(timezone string) error {
	_, err := time.LoadLocation(timezone)
	if err != nil {
		return fmt.Errorf("invalid timezone: %s", timezone)
	}
	return nil
}

// AddOneDay adds one day to a date string in YYYY-MM-DD format
func AddOneDay(dateStr string) (string, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", err
	}
	nextDay := date.AddDate(0, 0, 1)
	return nextDay.Format("2006-01-02"), nil
}
