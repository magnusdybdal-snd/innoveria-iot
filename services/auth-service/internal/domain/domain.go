// Package domain defines auth-service core interfaces and domain contracts.
package domain

import (
	"context"
)

// AuthService defines authentication and session operations exposed by
// the auth domain service layer.
type AuthService interface {
	// TODO: add these methods
	// Login
	// Refresh
	// Logout
	// Me
	RegisterCompany(ctx context.Context, payload Company) (Company, error)
	GetOneCompany(ctx context.Context, companyID string) (Company, error)
	GetAllCompanies(ctx context.Context) ([]Company, error)
}
