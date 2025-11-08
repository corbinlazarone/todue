package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.Handler(http.MethodGet, "/", http.HandlerFunc(app.health))

	// auth routes
	router.Handler(http.MethodPost, "/auth/login", http.HandlerFunc(app.loginHandler))

	std := alice.New(secureHeaders)

	return std.Then(router)
}
