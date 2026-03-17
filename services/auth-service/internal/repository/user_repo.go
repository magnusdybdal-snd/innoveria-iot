// Package repository contains auth-service persistence implementation.
package repository

import (
	"context"
	"innoveria-iot/auth-service/internal/domain"
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

// FindByEmail retrieves one user by email.
func (r *UserRepoImpl) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	return domain.User{}, nil
}

// UpdateLastLoggedIn updates the user's last login timestamp.
func (r *UserRepoImpl) UpdateLastLoggedIn(ctx context.Context, userID string) error {
	return nil
}
