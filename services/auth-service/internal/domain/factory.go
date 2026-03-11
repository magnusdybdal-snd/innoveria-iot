package domain

import (
	"context"
	"errors"
)

var (
	// ErrCompanyNotFound triggers when there is a sql violation for company id in factory
	ErrCompanyNotFound = errors.New("company not found")
	// ErrInvalidInput trigger when there is an invalid input for factory
	ErrInvalidInput = errors.New("invalid input")
)

// Factory is the domain model for factory
type Factory struct {
	ID        string
	CompanyID string
	Name      string
	Address   string
	CreatedAt string
	UpdatedAt string
}

// FactoryRepo defines the factory repository
type FactoryRepo interface {
	Create(ctx context.Context, factory Factory) (Factory, error)
	FindAll(ctx context.Context) ([]Factory, error)
	FindByID(ctx context.Context, factoryID string) (Factory, error)
	DeleteByID(ctx context.Context, factoryID string) error
}
