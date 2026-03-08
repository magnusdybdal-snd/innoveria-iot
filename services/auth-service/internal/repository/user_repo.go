// Package repository contains auth-service persistence implementation.
package repository

import (
	"context"
	"innoveria-iot/pkg/dbutil"
)

// UserRepoImpl is the PostgreSQL-backed implementation of
// auth domain user persistence operations.
type UserRepoImpl struct {
	db *dbutil.DB
}

// NewUserRepo initilize a new user repository
func NewUserRepo(db *dbutil.DB) *UserRepoImpl {
	return &UserRepoImpl{
		db: db,
	}
}

// Create inserts a new user in the database
func (r *UserRepoImpl) Create(ctx context.Context) error {
	return nil
}
