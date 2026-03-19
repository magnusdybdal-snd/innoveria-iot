// Package repository contains auth-service persistence implementation.
package repository

import (
	"context"
	"errors"
	"fmt"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5"
)

const (
	findUserByEmailQuery = `
		SELECT user_id, company_id, name, email, password_hash, role, last_logged_in, created_at, updated_at
		FROM auth."user"
		WHERE email = $1
	`
	findUserByIDQuery = `
		SELECT user_id, company_id, name, email, password_hash, role, last_logged_in, created_at, updated_at
		FROM auth."user"
		WHERE user_id = $1
	`
	updateLastLoggedInQuery = `
		UPDATE auth."user"
		SET last_logged_in = NOW()
		WHERE user_id = $1
	`
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
// Used by /login
func (r *UserRepoImpl) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var out domain.User
	err := r.db.Pool.QueryRow(ctx, findUserByEmailQuery,
		email,
	).Scan(
		&out.ID,
		&out.CompanyID,
		&out.Name,
		&out.Email,
		&out.PasswordHash,
		&out.Role,
		&out.LastLoggedIn,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// domain not found, code: 404
			return domain.User{}, fmt.Errorf("find user by email: %w", domain.ErrUserNotFound)
		}
		// Internal server error, code 500
		return domain.User{}, fmt.Errorf("find user by email: %w", err)
	}
	return out, nil
}

// FindByID retrieves one user by email.
// Used by /me
func (r *UserRepoImpl) FindByID(ctx context.Context, userID string) (domain.User, error) {
	var out domain.User
	err := r.db.Pool.QueryRow(ctx, userID,
		userID,
	).Scan(
		&out.ID,
		&out.CompanyID,
		&out.Name,
		&out.Email,
		&out.PasswordHash,
		&out.Role,
		&out.LastLoggedIn,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// domain not found, code: 404
			return domain.User{}, fmt.Errorf("find user by ID: %w", domain.ErrUserNotFound)
		}
		// Internal server error, code 500
		return domain.User{}, fmt.Errorf("find user by ID: %w", err)
	}
	return out, nil
}

// UpdateLastLoggedIn updates the user's last login timestamp.
func (r *UserRepoImpl) UpdateLastLoggedIn(ctx context.Context, userID string) error {
	res, err := r.db.Pool.Exec(ctx, updateLastLoggedInQuery, userID)
	if err != nil {
		return fmt.Errorf("update user last login: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("update user last login: %w", domain.ErrUserNotFound)
	}

	return nil
}
