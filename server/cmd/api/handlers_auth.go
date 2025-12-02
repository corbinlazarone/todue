package main

import (
	"encoding/json"
	"net/http"

	"github.com/corbinlazarone/cmovie/cmd/internals/auth"
	"github.com/corbinlazarone/cmovie/cmd/internals/models"
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

	tokenString, err := generateJWT(user)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Authentication failed")
		return
	}

	rep.WriteSuccessResponse(w, tokenString, http.StatusOK)
}
