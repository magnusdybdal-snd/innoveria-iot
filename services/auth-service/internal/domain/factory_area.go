package domain

import (
	"context"
	"time"
)

// FactoryArea represents a sub-unit of a factory (company -> factory -> factory_area)
type FactoryArea struct {
	ID          string
	FactoryID   string
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// FactoryAreaRepo defines the repository layer for factory area
type FactoryAreaRepo interface {
	Create(ctx context.Context, area FactoryArea) (FactoryArea, error)
	FindAll(ctx context.Context) ([]FactoryArea, error)
	FindByID(ctx context.Context, areaID string) (FactoryArea, error)
	Delete(ctx context.Context, areaID string) error
}

// FactoryAreaService defines factory area use-cases exposed by the service layer.
type FactoryAreaService interface {
	RegisterFactoryArea(ctx context.Context, payload FactoryArea) (FactoryArea, error)
	GetAllFactoryAreas(ctx context.Context) ([]FactoryArea, error)
	GetOneFactoryArea(ctx context.Context, areaID string) (FactoryArea, error)
	DeleteFactoryArea(ctx context.Context, areaID string) error
}
