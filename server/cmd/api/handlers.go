package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/corbinlazarone/cmovie/cmd/internals/models"
)

type Response struct {
	Message string `json:"message"`
}

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "cmovie server is running!")
}

func (app *application) CreateReview(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // close the request body when the function returns

	var review models.Review

	// NOTE: decode request body into review struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&review)
	if err != nil {
		// TODO: log error -- make some app level logger for this
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// TODO: validate review struct

	response := Response{
		Message: "Review created successfully!",
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		// TODO: log error -- make some app level logger for this
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBytes)
}
