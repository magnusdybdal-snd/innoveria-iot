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
	// ROLE_USER is the standard user role.
	ROLE_USER RoleType = "FACTORY_WORKER"
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
	// Create inserts a new user into the database.
	Create(ctx context.Context, user User) (User, error)
	// FindAll retrieves all users. If companyID is non-empty, results are filtered to that company.
	FindAll(ctx context.Context, companyID string) ([]User, error)
	// Update updates the editable fields of a user.
	Update(ctx context.Context, userID string, payload User) error
	// Delete removes a user by ID.
	Delete(ctx context.Context, userID string) error
}

// UserService defines the business logic operations for user management.
type UserService interface {
	// Create creates a new user, hashing the password before storage.
	Create(ctx context.Context, user User) (User, error)
	// GetAll retrieves all users. If companyID is non-empty, results are filtered to that company.
	GetAll(ctx context.Context, companyID string) ([]User, error)
	// Update updates the editable fields of a user.
	Update(ctx context.Context, userID string, payload User) error
	// Delete removes a user by ID.
	Delete(ctx context.Context, userID string) error
}
