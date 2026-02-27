package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/corbinlazarone/todue/cmd/internals/auth"
	"github.com/corbinlazarone/todue/cmd/internals/types"
)

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	defer r.Body.Close()

	// Limit request body size to 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var rep types.Response

	type parameters struct {
		AuthCode string `json:"authCode"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Invalid request body")
		return
	}

	if params.AuthCode == "" {
		rep.WriteErrorResponse(w, http.StatusBadRequest, "Authorization code is required")
		return
	}

	// Exchanged auth code for access and refresh token
	tokens, err := auth.ExchangeCodeForTokens(ctx, params.AuthCode)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Authentication failed")
		return
	}

	// Get user information from google userinfo endpoint
	userInfo, err := auth.GetUserInfo(tokens.AccessToken)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Authentication failed")
		return
	}

	// Create or update user if already exists
	user, err := app.users.CheckIfExists(ctx, userInfo)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Authentication failed")
		return
	}

	// if their is a refresh token present add it to user entry
	if tokens.RefreshToken != "" {
		err := app.users.UpdateRefreshToken(ctx, user.Email, tokens.RefreshToken)
		if err != nil {
			app.errLog.Println(err)
			rep.WriteErrorResponse(w, http.StatusInternalServerError, "Authentication failed")
			return
		}
	}

	// Generate todue JWT for client
	tokenString, err := generateJWT(user)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Authentication failed")
		return
	}

	rep.WriteSuccessResponse(w, tokenString, http.StatusOK)
}
