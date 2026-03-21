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
	Delete(ctx context.Context, factoryID string) error
}
