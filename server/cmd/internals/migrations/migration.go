package migrations

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

// NOTE: embeds migrations in to executable binary when we build go
// --- prettyy coool

//go:embed *.sql
var embedMigrationFS embed.FS

func RunMigrations(db *sql.DB) error {
	goose.SetBaseFS(embedMigrationFS)

	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	err = goose.Up(db, ".")
	if err != nil {
		return err
	}

	return nil
}
