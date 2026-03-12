package db

import (
	"embed"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed seeds/*.sql
var seeds embed.FS

// RunSeeds see /pkg/dbutil/seed.go
func RunSeeds(pool *pgxpool.Pool) error {
	return dbutil.RunSeeds(pool, seeds)
}
