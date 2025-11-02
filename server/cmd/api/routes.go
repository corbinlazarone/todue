package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.Handler(http.MethodGet, "/", http.HandlerFunc(app.health))
	router.Handler(http.MethodPost, "/reviews/create", http.HandlerFunc(app.CreateReview))

	std := alice.New(secureHeaders)

	return std.Then(router)
}
