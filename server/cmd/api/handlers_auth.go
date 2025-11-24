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

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Limit request body size to 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

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
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Authentication failed")
		return
	}

	tokenString, err := app.generateJWT(user)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Authentication failed")
		return
	}

	rep.WriteSuccessResponse(w, tokenString, http.StatusOK)
}

func (app *application) generateJWT(user models.User) (string, error) {
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
		return "", errors.New("JWT_SECRET not set")
	}

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
