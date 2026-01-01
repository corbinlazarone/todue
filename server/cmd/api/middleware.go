package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/corbinlazarone/todue/cmd/internals/types"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

// add secure headers to all incoming request
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")

		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		// NOTE: Any code here will execute on the way down the chain.
		// Any early returns will break the chain.

		next.ServeHTTP(w, r)

		// NOTE: Any code here will execute on the way back up the chain.
	})
}

// enableCORS adds CORS headers to allow requests from the Next.js client
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		allowedOrigin := os.Getenv("WEB_CLIENT_URL")
		if allowedOrigin == "" {
			log.Fatal("WEB_CLIENT_URL env var not set")
		}

		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Show 500 internal server error to user on request panics
func (app *application) recoverFromPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// defered function that will alwasy run in the event of a panic as Go unwinds
		// the stack.
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connetion", "close")

				// show a 500 server error to the user
				var rep types.Response
				app.errLog.Println(err)
				rep.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// Allow 2 requests per second with burst of 10 for AI extraction
var limiter = rate.NewLimiter(2, 10)

func rateLimter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var rep types.Response

		// get bearer token from header
		authHeader := r.Header.Get("Authorization")
		extractedToken, err := extractBearerToken(authHeader)
		if err != nil {
			rep.WriteErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			log.Fatal("JWT_SECRET not set")
		}

		// validate token using golang-jwt/jwt package
		token, err := jwt.Parse(extractedToken, func(token *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil || !token.Valid {
			rep.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		userID, err := token.Claims.GetSubject()
		if err != nil {
			log.Print("WARNGING: Couldn't get user id from token")
			rep.WriteErrorResponse(w, http.StatusUnauthorized, "Internal Server Error")
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)

		// token is valid - continue with the request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
