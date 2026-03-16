// Package db embeds and runs SQL migrations and seeds for the device service database.
package db

import (
	"embed"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Bundles all sql files in the migrations folder directly into the compiled binary
// Declared here because //go:embed resolves relative to this file's location at compile time.
//
//go:embed migrations/*.sql
var migrations embed.FS

// RunMigrations runs all embedded SQL migration files against the given database pool.
func RunMigrations(pool *pgxpool.Pool) error {
	return dbutil.RunMigrations(pool, migrations)
}
