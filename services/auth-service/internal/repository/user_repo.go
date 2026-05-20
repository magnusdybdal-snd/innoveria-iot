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
		SELECT user_id, company_id, name, email, role, last_logged_in, created_at, updated_at
		FROM auth."user"
		WHERE user_id = $1
	`
	updateLastLoggedInQuery = `
		UPDATE auth."user"
		SET last_logged_in = NOW()
		WHERE user_id = $1
	`

	createUserQuery = `
		INSERT INTO auth."user" (company_id, name, email, password_hash, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id, company_id, name, email, role, last_logged_in, created_at, updated_at
	`

	findAllUsersQuery = `
		SELECT user_id, company_id, name, email, role, last_logged_in, created_at, updated_at
		FROM auth."user"
		ORDER BY created_at ASC
	`

	findAllUsersByCompanyQuery = `
		SELECT user_id, company_id, name, email, role, last_logged_in, created_at, updated_at
		FROM auth."user"
		WHERE company_id = $1
		ORDER BY created_at ASC
	`

	updateUserQuery = `
		UPDATE auth."user"
		SET name = $2, email = $3, password_hash = $4, updated_at = NOW()
		WHERE user_id = $1
	`

	deleteUserQuery = `
		DELETE FROM auth."user"
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
	err := r.db.Pool.QueryRow(ctx, findUserByIDQuery,
		userID,
	).Scan(
		&out.ID,
		&out.CompanyID,
		&out.Name,
		&out.Email,
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

// Create inserts a new user into the database.
func (r *UserRepoImpl) Create(ctx context.Context, user domain.User) (domain.User, error) {
	var out domain.User
	err := r.db.Pool.QueryRow(ctx, createUserQuery,
		user.CompanyID,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Role,
	).Scan(
		&out.ID,
		&out.CompanyID,
		&out.Name,
		&out.Email,
		&out.Role,
		&out.LastLoggedIn,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}
	return out, nil
}

// FindAll retrieves all users. If companyID is non-empty, results are filtered to that company.
func (r *UserRepoImpl) FindAll(ctx context.Context, companyID string) ([]domain.User, error) {
	var rows pgx.Rows
	var err error

	if companyID != "" {
		rows, err = r.db.Pool.Query(ctx, findAllUsersByCompanyQuery, companyID)
	} else {
		rows, err = r.db.Pool.Query(ctx, findAllUsersQuery)
	}
	if err != nil {
		return nil, fmt.Errorf("find all users: %w", err)
	}
	defer rows.Close()

	users := []domain.User{}
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(
			&u.ID,
			&u.CompanyID,
			&u.Name,
			&u.Email,
			&u.Role,
			&u.LastLoggedIn,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("find all users: %w", err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("find all users: %w", err)
	}

	return users, nil
}

// Update updates the editable fields of a user.
func (r *UserRepoImpl) Update(ctx context.Context, userID string, payload domain.User) error {
	res, err := r.db.Pool.Exec(ctx, updateUserQuery,
		userID,
		payload.Name,
		payload.Email,
		payload.PasswordHash,
	)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("update user: %w", domain.ErrUserNotFound)
	}

	return nil
}

// Delete removes a user by ID.
func (r *UserRepoImpl) Delete(ctx context.Context, userID string) error {
	res, err := r.db.Pool.Exec(ctx, deleteUserQuery, userID)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("delete user: %w", domain.ErrUserNotFound)
	}

	return nil
}
