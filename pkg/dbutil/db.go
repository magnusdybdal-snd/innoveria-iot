// Package dbutil provides shared PostgreSQL utilities for connecting to
// and managing databases across services. It handles connection pooling,
// schema migrations via goose, and environment-aware seeding.
package dbutil

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB wraps a pgxpool connection pool for safe concurrent database access.
type DB struct {
	Pool *pgxpool.Pool
}

// New creates a new DB connection pool using the provided connection string.
// serviceName is used in the startup log to identify which service connected.
func New(connString string, serviceName string) (*DB, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	// Config for handling the number of concurrent connection to the postgres pool
	cfg.MaxConns = 10 // connection after this needs to wait
	cfg.MinConns = 0  // no idle connections kept alive, pool creates on demand
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	// Pinging the database
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	slog.Info(serviceName + " successfully connected")

	return &DB{Pool: pool}, nil
}

// Close releases all connections in the pool.
func (d *DB) Close() {
	d.Pool.Close()
}
