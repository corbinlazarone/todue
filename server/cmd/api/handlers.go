package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/corbinlazarone/cmovie/cmd/internals/models"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "cmovie server is running!")
}

func (app *application) CreateReview(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // close the request body when the function returns

	var review models.Review
	var response models.Response

	// NOTE: decode request body into review struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&review)
	if err != nil {
		app.errLog.Println(err)
		response.WriteErrorResponse(w, "Invalid JSON")
		return
	}

	// TODO: validate review struct

	// NOTE: send success response
	response.WriteSuccessResponse(w, "Review created successfully!", http.StatusCreated)
}
