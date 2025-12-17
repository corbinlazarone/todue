package main

import (
	"net/http"

	"github.com/corbinlazarone/Todue-Actual/cmd/internals/types"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	var response types.Response
	response.WriteSuccessResponse(w, "Todue server is running..", http.StatusOK)
}
