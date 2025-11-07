package main

import (
	"net/http"

	"github.com/corbinlazarone/cmovie/cmd/internals/models"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	response.WriteSuccessResponse(w, "cmovie server is running..", http.StatusOK)
}
