package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/corbinlazarone/cmovie/cmd/internals/auth"
	"github.com/corbinlazarone/cmovie/cmd/internals/models"
	"github.com/golang-jwt/jwt/v5"
)

// TODO: add rate limiting for this endpoint
func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var rep models.Response

	type parameters struct {
		GoogleJWT string `json:"googleJWT"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Invalid request body")
		return
	}

	// validate google jwt
	claims, err := auth.ValidateGoogleJWT(params.GoogleJWT)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Google JWT")
		return
	}

	user, err := app.users.CheckIfExists(claims)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Error checking if user exists")
		return
	}

	tokenString, err := generateJWT(user, claims)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Error generating JWT")
		return
	}

	rep.WriteSuccessResponse(w, tokenString, http.StatusOK)
}

func generateJWT(user models.User, googleClaims auth.GoogleClaims) (string, error) {
	claims := jwt.MapClaims{
		"sub":        user.Id,
		"iat":        time.Now().Unix(),
		"exp":        googleClaims.ExpiresAt.Unix(), // Expires same time as Google JWT
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
		return "", errors.New("JWT_SECRET not set")
	}

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
