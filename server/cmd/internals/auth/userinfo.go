package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GoogleUserInfo struct {
	Sub           string `json:"sub"` // Google's unique user ID
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	Picture       string `json:"picture"`
}

func GetUserInfo(accessToken string) (GoogleUserInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", accessToken))
	if err != nil {
		return GoogleUserInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GoogleUserInfo{}, fmt.Errorf("failed to get user info: status %d", resp.StatusCode)
	}

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return GoogleUserInfo{}, err
	}

	return userInfo, nil
}
