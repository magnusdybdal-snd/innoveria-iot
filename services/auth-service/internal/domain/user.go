package domain

import (
	"context"
	"time"
)

// RoleType represents the authorization role assigned to a user.
type RoleType string

const (
	// ROLE_PLATFORM_ADMIN is the platform administrator role.
	ROLE_PLATFORM_ADMIN RoleType = "PLATFORM_ADMIN"
	// TODO: add the rest of roles
)

// User is the domain model for auth users.
type User struct {
	ID           string
	CompanyID    string
	Name         string
	Email        string
	PasswordHash string
	Role         RoleType
	LastLoggedIn *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// UserRepo defines user persistence operations required by auth workflows.
type UserRepo interface {
	// FindByEmail retrieves one user by email.
	FindByEmail(ctx context.Context, email string) (User, error)
	// FindByID retrieves one user by ID.
	FindByID(ctx context.Context, userID string) (User, error)
	// UpdateLastLoggedIn updates the user's last login timestamp.
	UpdateLastLoggedIn(ctx context.Context, userID string) error
}
