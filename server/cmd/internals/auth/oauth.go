package auth

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func OAuth2Config() *oauth2.Config {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		log.Fatal("GOOGLE_CLIENT_ID is required")
	}

	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("GOOGLE_CLIENT_ID is required")
	}

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "postmessage",
		Scopes: []string{
			"https://www.googleapis.com/auth/calendar",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

// ExchangeCodeForTokens exchanges authorization code for access and refresh tokens
func ExchangeCodeForTokens(ctx context.Context, code string) (*oauth2.Token, error) {
	config := OAuth2Config()
	token, err := config.Exchange(ctx, code)
	return token, err
}

func EncryptRefreshToken(plainText string) (string, error) {
	key := []byte(os.Getenv("TOKEN_ENCRYPTION_KEY"))
	if len(key) != 32 {
		log.Fatalf("Token Encryption key must be 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptRefreshToken(encrypted string) (string, error) {
	key := []byte(os.Getenv("TOKEN_ENCRYPTION_KEY"))
	if len(key) != 32 {
		log.Fatalf("Token Encryption key must be 32 bytes")
	}

	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce := data[:nonceSize]
	ciphertextBytes := data[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
