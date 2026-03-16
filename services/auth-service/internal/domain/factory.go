package domain

import (
	"context"
	"time"
)

// Factory is the domain model for factory
type Factory struct {
	ID        string
	CompanyID string
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FactoryRepo defines the factory repository
type FactoryRepo interface {
	Create(ctx context.Context, factory Factory) (Factory, error)
	FindAll(ctx context.Context) ([]Factory, error)
	FindByID(ctx context.Context, factoryID string) (Factory, error)
	DeleteByID(ctx context.Context, factoryID string) error
}

// FactoryService defines factory use-cases exposed by the service layer.
type FactoryService interface {
	RegisterFactory(ctx context.Context, payload Factory) (Factory, error)
	GetOneFactory(ctx context.Context, factoryID string) (Factory, error)
	GetAllFactories(ctx context.Context) ([]Factory, error)
	DeleteFactory(ctx context.Context, factoryID string) error
}
