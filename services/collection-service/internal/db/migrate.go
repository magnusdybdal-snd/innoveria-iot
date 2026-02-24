package db

import (
	"embed"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// Bundles all sql files in the migrations folder directly into the compiled binary
//
//go:embed migrations/*.sql
var migrations embed.FS

func RunMigrations(pool *pgxpool.Pool) error {
	// Convert the pool to sqldb for goose to work on it
	sqlDB := stdlib.OpenDBFromPool(pool)

	// Teels goose to read from the embeded files instead of reading from disk
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	// Runs all pending migrations. If already applied, goose skips them
	return goose.Up(sqlDB, "migrations")
}
