package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// public routes
	router.Handler(http.MethodGet, "/", http.HandlerFunc(app.health))
	router.Handler(http.MethodPost, "/api/auth/login", http.HandlerFunc(app.loginHandler))

	// Protected routes
	protected := alice.New(requireAuth)
	router.Handler(http.MethodPost, "/api/ai/extract", protected.ThenFunc(app.extractCourseData))
	router.Handler(http.MethodPost, "/api/event/save", http.HandlerFunc(app.insertCourseDataHandler))

	std := alice.New(enableCORS, secureHeaders, rateLimter, app.recoverFromPanic)

	return std.Then(router)
}
