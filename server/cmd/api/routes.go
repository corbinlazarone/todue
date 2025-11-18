package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	limter := alice.New(rateLimter)

	router.Handler(http.MethodGet, "/", http.HandlerFunc(app.health))
	router.Handler(http.MethodPost, "/auth/login", limter.ThenFunc(app.loginHandler))

	// Protected routes
	router.Handler(http.MethodPost, "/ai/extract", http.HandlerFunc(app.ExtractCourseData))

	std := alice.New(enableCORS, secureHeaders)

	return std.Then(router)
}
