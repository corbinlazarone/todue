package models

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
}

func validateGoogleJWT(token string) (GoogleClaims, error) {
	return GoogleClaims{}, nil
}
