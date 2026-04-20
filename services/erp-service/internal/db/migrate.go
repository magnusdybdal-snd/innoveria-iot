// Package db contains database migration logic for the erp service.
package db

import (
	"embed"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.sql
var migrations embed.FS

// RunMigrations runs goose migrations see /pkg/dbutil/migrate.go
func RunMigrations(pool *pgxpool.Pool) error {
	return dbutil.RunMigrations(pool, migrations)
}
