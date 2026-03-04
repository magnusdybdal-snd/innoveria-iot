// Package db TODO(@vinjar): add proper documentation.
package db

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB TODO(@vinjar): add proper documentation.
type DB struct {
	Pool *pgxpool.Pool
}

// New uses pgxpool instead of pgx, for threadsafe database writing
func New(connString string) (*DB, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	// Config for handling the number of concurrent connection to the postgres pool
	cfg.MaxConns = 20 // connection after this needs to wait
	cfg.MinConns = 5  // keeps 5 connections open at all times
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	// Pinging the time series database
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	slog.Info("Timescale DB successfully connected")

	return &DB{Pool: pool}, nil
}

// Close TODO(@vinjar): add proper documentation.
func (d *DB) Close() {
	d.Pool.Close()
}
