// Package domain defines auth-service core interfaces and domain contracts.
package domain

import (
	"context"
	"time"
)

// CompanyRepo defines the company repository needed by the auth domain.
type CompanyRepo interface {
	Create(ctx context.Context, company Company) (Company, error)
	// TODO:
	// Get all
	// Get one
	// Put/patch
}

// AuthService defines authentication and session operations exposed by
// the auth domain service layer.
type AuthService interface {
	// TODO: add these methods
	// Login
	// Refresh
	// Logout
	// Me
	RegisterCompany(ctx context.Context, payload Company) (Company, error)
}

// Company is the domain model for company
type Company struct {
	CompanyID string
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
