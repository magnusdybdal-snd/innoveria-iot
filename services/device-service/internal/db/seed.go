package db

import (
	"embed"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Bundles all sql files in the seeds folder directly into the compiled binary
// Declared here because //go:embed resolves relative to this file's location at compile time.
//
//go:embed seeds/*.sql
var seeds embed.FS

// RunSeeds TODO(@Magnus Dybdal): add proper documentation.
func RunSeeds(pool *pgxpool.Pool) error {
	return dbutil.RunSeeds(pool, seeds)
}
