package db

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(connString string) (*DB, error) {
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

	slog.Info("Device DB successfully connected")

	return &DB{Pool: pool}, nil
}

func (d *DB) Close() {
	d.Pool.Close()
}
