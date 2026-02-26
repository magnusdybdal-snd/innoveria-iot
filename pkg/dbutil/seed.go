package dbutil

import (
	"embed"
	"fmt"
	"innoveria-iot/pkg/env"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

// RunSeeds executes a seed file matching the current GO_ENV environment variable.
// It expects seeds to contain a file at "seeds/<env>.sql" (e.g. seeds/development.sql).
// No-ops in production. If no matching seed file is found, it logs a warning and skips.
func RunSeeds(pool *pgxpool.Pool, seeds embed.FS) error {
	// Check if running production or development
	envVar := env.Get("GO_ENV", "development")

	if envVar == "production" {
		return nil
	}

	fileName := fmt.Sprintf("seeds/%s.sql", envVar)

	// Extract the seed. If not found: warning and skip
	// since some db in dev might not need mock data
	seed, err := seeds.ReadFile(fileName)
	if err != nil {
		slog.Warn("no seed file found, skipping", "environment", envVar)
		return nil
	}

	// Convert from pool to sqldb for Exec
	sqlDB := stdlib.OpenDBFromPool(pool)

	// Executes the seed (inserts mock data)
	_, err = sqlDB.Exec(string(seed))
	return err
}
