// Package db contains database migration and seed logic for the erp service.
package db

import (
	"embed"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed seeds/*.sql
var seeds embed.FS

// RunSeeds runs all embedded SQL seed files against the given database pool.
func RunSeeds(pool *pgxpool.Pool) error {
	return dbutil.RunSeeds(pool, seeds)
}
