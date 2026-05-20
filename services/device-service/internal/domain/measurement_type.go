package domain

import "context"

// MeasurementType is the canonical vocabulary entry for a sensor measurement.
// The slug is immutable once created and acts as the cross-service contract.
// Types are never deleted — use Deprecate to retire a type.
type MeasurementType struct {
	Slug        string
	DisplayName string
	Description *string
	DefaultUnit *string
	Deprecated  bool
}

// MeasurementTypeRepository handles persistence of measurement types.
type MeasurementTypeRepository interface {
	Create(ctx context.Context, m MeasurementType) error
	FindActive(ctx context.Context) ([]MeasurementType, error)
	FindAll(ctx context.Context) ([]MeasurementType, error)
	Deprecate(ctx context.Context, slug string) error
}

// MeasurementTypeService defines the business logic for managing measurement types.
type MeasurementTypeService interface {
	Create(ctx context.Context, m MeasurementType) error
	ListActive(ctx context.Context) ([]MeasurementType, error)
	ListAll(ctx context.Context) ([]MeasurementType, error)
	Deprecate(ctx context.Context, slug string) error
}
