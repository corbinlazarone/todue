package models

import (
	"context"
	"errors"
	"time"

	"github.com/corbinlazarone/Todue-Actual/cmd/internals/auth"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Id            string    `json:"id"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"emailVerified" db:"email_verified"`
	FirstName     string    `json:"firstName" db:"first_name"`
	LastName      string    `json:"lastName" db:"last_name"`
	Picture       string    `json:"picture"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
}

type UserModel struct {
	DB *pgxpool.Pool
}

// CheckIfExists checks if a user with the given googleID or email exists in the database.
// If not, it creates a new user.
// If user exists by email but has different google_id, it updates the google_id.
func (u *UserModel) CheckIfExists(claims auth.GoogleClaims) (User, error) {
	statement := `SELECT id, email, email_verified, first_name, last_name, picture, created_at, updated_at
                  FROM users
                  WHERE email = $1`

	rows, err := u.DB.Query(context.Background(), statement, claims.Email)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// User doesn't exist - create new one
			newUser, err := u.createNewUser(claims)
			if err != nil {
				return User{}, err
			}
			return newUser, nil
		}
		return User{}, err
	}

	return user, nil
}

func (u *UserModel) createNewUser(claims auth.GoogleClaims) (User, error) {
	statement := `INSERT INTO users (email, email_verified, first_name, last_name, picture)
                VALUES ($1, $2, $3, $4, $5)
                RETURNING id, email, email_verified, first_name, last_name, picture, created_at, updated_at`

	var user User
	rows, err := u.DB.Query(context.Background(), statement,
		claims.Email,
		claims.EmailVerified,
		claims.FirstName,
		claims.LastName,
		claims.Picture,
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
