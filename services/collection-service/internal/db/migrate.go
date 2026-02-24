package db

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

// Bundles all sql files in the migrations folder directly into the compiled binary
//
//go:embed migrations/*.sql
var migrations embed.FS

func RunMigrations(db *sql.DB) error {
	// Teels goose to read from the embeded files instead of reading from disk
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	// Runs all pending migrations. If already applied, goose skips them
	return goose.Up(db, "migrations")
}
