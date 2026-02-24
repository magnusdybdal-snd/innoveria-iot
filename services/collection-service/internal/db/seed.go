package db

import (
	"embed"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

//go:embed seeds/*.sql
var seeds embed.FS

func RunSeeds(pool *pgxpool.Pool) error {
	// Check if running production or development
	env := os.Getenv("GO_ENV")

	if env == "production" {
		return nil
	}

	fileName := fmt.Sprintf("seeds/%s.sql", env)

	// Extract the seed. If not found: warning and skip
	// since some db in dev might not need mock data
	seed, err := seeds.ReadFile(fileName)
	if err != nil {
		slog.Warn("no seed file found, skipping", "environment", env)
		return nil
	}

	// Convert from pool to sqldb for goose to work on
	sqlDB := stdlib.OpenDBFromPool(pool)

	// Executes the seed (inserts mock data)
	_, err = sqlDB.Exec(string(seed))
	return err
}
