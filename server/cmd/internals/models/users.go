package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/corbinlazarone/Todue-Actual/cmd/internals/auth"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Id                 string    `json:"id"`
	Email              string    `json:"email"`
	EmailVerified      bool      `json:"emailVerified" db:"email_verified"`
	FirstName          string    `json:"firstName" db:"first_name"`
	LastName           string    `json:"lastName" db:"last_name"`
	Picture            string    `json:"picture"`
	CreatedAt          time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt          time.Time `json:"updatedAt" db:"updated_at"`
	GoogleRefreshToken string    `json:"google_refresh_token"`
}

type UserModel struct {
	DB *pgxpool.Pool
}

// CheckIfExists checks if a user with the given googleID or email exists in the database.
// If not, it creates a new user.
// If user exists by email but has different google_id, it updates the google_id.
func (u *UserModel) CheckIfExists(ctx context.Context, userInfo auth.GoogleUserInfo) (User, error) {
	statement := `SELECT id, email, email_verified, first_name, last_name, picture, created_at, updated_at 
                  FROM users
                  WHERE email = $1`

	rows, err := u.DB.Query(ctx, statement, userInfo.Email)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// User doesn't exist - create new one
			newUser, err := u.createNewUser(ctx, userInfo)
			if err != nil {
				return User{}, err
			}
			return newUser, nil
		}
		return User{}, err
	}

	return user, nil
}

func (u *UserModel) GetByID(ctx context.Context, userID string) (*User, error) {
	statement := `SELECT id, email, first_name, last_name, picture, created_at, 
	updated_at, google_refresh_token FROM users WHERE id = $1`

	var user User
	err := u.DB.QueryRow(ctx, statement, userID).Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Picture,
		&user.GoogleRefreshToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserModel) UpdateRefreshToken(ctx context.Context, email string, refreshToken string) error {
	encryptedToken, err := auth.EncryptRefreshToken(refreshToken)
	if err != nil {
		return fmt.Errorf("failed to encrypt refresh token: %w", err)
	}

	statement := `UPDATE users SET google_refresh_token = $1, updated_at = NOW() WHERE email = $2`
	_, err = u.DB.Exec(ctx, statement, encryptedToken, email)
	return err
}

func (u *UserModel) createNewUser(ctx context.Context, userInfo auth.GoogleUserInfo) (User, error) {
	statement := `INSERT INTO users (email, email_verified, first_name, last_name, picture)
                VALUES ($1, $2, $3, $4, $5)
                RETURNING id, email, email_verified, first_name, last_name, picture, created_at, updated_at`

	var user User
	rows, err := u.DB.Query(ctx, statement,
		userInfo.Email,
		userInfo.EmailVerified,
		userInfo.FirstName,
		userInfo.LastName,
		userInfo.Picture,
	)
	if err != nil {
		return User{}, err
	}

	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return User{}, err
	}

	return user, nil
}
