package db

import (
	"embed"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed seeds/*.sql
var seeds embed.FS

// RunSeeds is a thin wrapper around dbutil.RunSeeds
func RunSeeds(pool *pgxpool.Pool) error {
	return dbutil.RunSeeds(pool, seeds)
}
