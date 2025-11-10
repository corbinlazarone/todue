package models

import (
	"context"
	"errors"
	"time"

	"github.com/corbinlazarone/cmovie/cmd/internals/auth"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Id            string    `json:"id"`
	GoogleID      string    `json:"googleID" db:"google_id"`
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
	statement := `SELECT id, google_id, email, email_verified, first_name, last_name, picture, created_at, updated_at
                  FROM users
                  WHERE google_id = $1 OR email = $2`

	rows, _ := u.DB.Query(context.Background(), statement, claims.Sub, claims.Email)
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

	// User exists - check if we need to update google_id
	if user.GoogleID != claims.Sub {
		updateStatement := `UPDATE users SET google_id = $1, updated_at = NOW() WHERE id = $2`
		_, err := u.DB.Exec(context.Background(), updateStatement, claims.Sub, user.Id)
		if err != nil {
			return User{}, err
		}
		user.GoogleID = claims.Sub
		user.UpdatedAt = time.Now()
	}

	return user, nil
}

func (u *UserModel) createNewUser(claims auth.GoogleClaims) (User, error) {
	statement := `INSERT INTO users (google_id, email, email_verified, first_name, last_name, picture)
                VALUES ($1, $2, $3, $4, $5, $6)
                RETURNING id, google_id, email, email_verified, first_name, last_name, picture, created_at, updated_at`

	var user User
	rows, _ := u.DB.Query(context.Background(), statement,
		claims.Sub,
		claims.Email,
		claims.EmailVerified,
		claims.FirstName,
		claims.LastName,
		claims.Picture,
	)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])

	if err != nil {
		return User{}, err
	}

	return user, nil
}
