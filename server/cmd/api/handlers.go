package main

import (
	"encoding/json"
	"net/http"

	"github.com/corbinlazarone/cmovie/cmd/internals/models"
	"github.com/corbinlazarone/cmovie/cmd/internals/validator"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	response.WriteSuccessResponse(w, "cmovie server is running..", http.StatusOK)
}

func (app *application) CreateReview(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // close the request body when the function returns

	var review models.Review
	var v validator.Validator
	var response models.Response

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&review)
	if err != nil {
		app.errLog.Println(err)
		response.WriteErrorResponse(w, "Invalid JSON")
		return
	}

	v.CheckField(validator.NotBlank(review.Title), "title", "Title cannot be blank")
	v.CheckField(validator.MinChars(review.Title, 5), "title", "Title must be at least 5 characters")
	v.CheckField(validator.MaxChars(review.Title, 100), "title", "Title cannot be more than 100 characters")
	if !v.Valid() {
		response.WriteErrorResponse(w, v.FieldErrors["title"])
		return
	}

	v.CheckField(validator.NotBlank(review.Summary), "summary", "Summary cannot be blank")
	v.CheckField(validator.MinChars(review.Summary, 10), "summary", "Summary must be at least 10 characters")
	v.CheckField(validator.MaxChars(review.Summary, 500), "summary", "Summary cannot be more than 500 characters")
	if !v.Valid() {
		response.WriteErrorResponse(w, v.FieldErrors["summary"])
		return
	}

	v.CheckField(validator.PermittedInt(review.Rating, 1, 2, 3, 4, 5), "rating", "Rating must be between 1 and 5")
	if !v.Valid() {
		response.WriteErrorResponse(w, v.FieldErrors["rating"])
		return
	}

	// TODO: write review to db and return id

	response.WriteSuccessResponse(w, "Review created successfully!", http.StatusCreated)
}
