package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GoogleClaims struct {
	Sub           string `json:"sub"` // Google's unique user ID
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	Picture       string `json:"picture"`
	jwt.RegisteredClaims
}

func ValidateGoogleJWT(tokenString string) (GoogleClaims, error) {
	claimsStuct := GoogleClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claimsStuct, func(token *jwt.Token) (any, error) {
		pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
		if err != nil {
			return nil, err
		}
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, err
		}
		return key, nil
	})

	if err != nil {
		return GoogleClaims{}, err
	}

	claims, ok := token.Claims.(*GoogleClaims)
	if !ok {
		return GoogleClaims{}, errors.New("Invalid Google JWT")
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return GoogleClaims{}, errors.New("iss is invalid")
	}

	if !slices.Contains(claims.Audience, os.Getenv("GOOGLE_CLIENT_ID")) {
		return GoogleClaims{}, errors.New("aud is invalid")
	}

	if claims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
		return GoogleClaims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}

// getGooglePublicKey returns the public key from google for the given keyID
// so we can verify a request.
func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")

	if err != nil {
		return "", err
	}

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	myResp := map[string]string{}

	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}

	key, ok := myResp[keyID]

	if !ok {
		return "", errors.New("key not found")
	}

	return key, nil
}
