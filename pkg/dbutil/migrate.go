package dbutil

import (
	"embed"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// RunMigrations applies all pending SQL migrations to the database using goose.
// migrations must be an embedded filesystem containing .sql files in a "migrations/" directory.
// Already applied migrations are skipped automatically.
func RunMigrations(pool *pgxpool.Pool, migrations embed.FS) error {
	// Convert the pool to sqldb for goose to work on it
	sqlDB := stdlib.OpenDBFromPool(pool)

	// Tells goose to read from the embedded files instead of reading from disk
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	// Runs all pending migrations. If already applied, goose skips them
	return goose.Up(sqlDB, "migrations")
}
