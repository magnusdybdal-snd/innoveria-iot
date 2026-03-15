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

	// Company
	RegisterCompany(ctx context.Context, payload Company) (Company, error)
	GetOneCompany(ctx context.Context, companyID string) (Company, error)
	GetAllCompanies(ctx context.Context) ([]Company, error)
	DeleteCompany(ctx context.Context, companyID string) error
	// Factory
	RegisterFactory(ctx context.Context, payload Factory) (Factory, error)
	GetOneFactory(ctx context.Context, factoryID string) (Factory, error)
	GetAllFactories(ctx context.Context) ([]Factory, error)
	DeleteFactory(ctx context.Context, factoryID string) error
	// Factory Area
	RegisterFactoryArea(ctx context.Context, payload FactoryArea) (FactoryArea, error)
	GetAllFactoryAreas(ctx context.Context) ([]FactoryArea, error)
	GetOneFactoryArea(ctx context.Context, areaID string) (FactoryArea, error)
}
