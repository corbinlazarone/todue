package main

import (
	"encoding/json"
	"net/http"

	"github.com/corbinlazarone/cmovie/cmd/internals/auth"
	"github.com/corbinlazarone/cmovie/cmd/internals/models"
)

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
	_, err = auth.ValidateGoogleJWT(params.GoogleJWT)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Google JWT")
		return
	}

	// TODO: Check if user exists in db, if not create user

	// TODO: create JWT token for our frontend and return it

	rep.WriteSuccessResponse(w, "Login successful", http.StatusOK)
}
