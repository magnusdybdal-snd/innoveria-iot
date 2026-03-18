// Package db provides database migration and seed utilities for context-service.
package db

import (
	"embed"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.sql
var migrations embed.FS

// RunMigrations is a thin wrapper around dbutil.RunMigrations
func RunMigrations(pool *pgxpool.Pool) error {
	return dbutil.RunMigrations(pool, migrations)
}
