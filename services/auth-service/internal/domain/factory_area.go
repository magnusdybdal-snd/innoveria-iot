package domain

import (
	"context"
	"time"
)

// FactoryArea is the domain model for factory area
type FactoryArea struct {
	ID          string
	FactoryID   string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// FactoryAreaRepo defines the repository layer for factory area
type FactoryAreaRepo interface {
	Create(ctx context.Context, area FactoryArea) (FactoryArea, error)
	FindAll(ctx context.Context) ([]FactoryArea, error)
	FindByID(ctx context.Context, areaID string) (FactoryArea, error)
	DeleteByID(ctx context.Context, areaID string) error
}
