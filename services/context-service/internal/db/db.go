// Package db provides database connection management for the context-service.
package db

import (
	"innoveria-iot/pkg/dbutil"
)

// DB is an alias for dbutil.DB
type DB = dbutil.DB

// New creates a new context-service database connection pool.
func New(connString string) (*DB, error) {
	return dbutil.New(connString, "context-db")
}
