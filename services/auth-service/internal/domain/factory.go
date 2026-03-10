package domain

import (
	"context"
)

// Factory is the domain model for factory
type Factory struct {
	Id         string
	CompanyID  string
	Name       string
	Address    string
	Created_at string
	Updated_at string
}

// FactoryRepo defines the factory repository
type FactoryRepo interface {
	Create(ctx context.Context, factory Factory) (Factory, error)
	FindAll(ctx context.Context) ([]Factory, error)
	FindByID(ctx context.Context, factoryID string) (Factory, error)
	DeleteByID(ctx context.Context, factoryID string) error
}
