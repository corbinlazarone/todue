package services

import (
	"context"
	"fmt"
	"log"

	"github.com/corbinlazarone/Todue-Actual/cmd/internals/auth"
	"github.com/corbinlazarone/Todue-Actual/cmd/internals/models"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// GoogleCalendarService() returns an authenticated Google Calendar service for the user.
// It handles token refresh automatically and saves new refresh tokens if issued.
func GoogleCalendarService(ctx context.Context, userModel *models.UserModel, userID, encryptedRFToken string) (*calendar.Service, error) {
	refreshToken, err := auth.DecryptRefreshToken(encryptedRFToken)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt refresh token: %w", err)
	}

	// Create OAuth2 config and token source
	config := auth.OAuth2Config()
	token := &oauth2.Token{RefreshToken: refreshToken}
	tokenSource := config.TokenSource(ctx, token)

	// Get current token
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get valid token: %w", err)
	}

	// If Google issued a new refresh token, save it
	if newToken.RefreshToken != "" && newToken.RefreshToken != refreshToken {
		if err := userModel.UpdateRefreshToken(ctx, userID, newToken.RefreshToken); err != nil {
			// Log but don't fail - the current token still works
			log.Printf("WARNING: failed to save new refresh token for user %s: %v\n", userID, err)
		}
	}

	// Create and return the Calendar service
	return calendar.NewService(ctx, option.WithTokenSource(tokenSource))
}
