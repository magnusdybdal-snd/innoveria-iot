package domain

import (
	"context"
	"errors"
	"time"
)

// ErrNotFound is returned when a requested resource does not exist.
var ErrNotFound = errors.New("not found")

// ContextData is the domain model for context data.
type ContextData struct {
	ID                   string
	CompanyID            string
	ContextType          string
	OrderID              string
	ProductionResourceID string
	DeviceEUI            string
	Value                float64
	Unit                 string
	PeriodStart          time.Time
	PeriodEnd            time.Time
	CalculatedAt         time.Time
}

// ContextService defines the context service interface for the context domain.
type ContextService interface {
	GetContextData(
		ctx context.Context,
		companyID,
		deviceEUI,
		contextType string,
		from,
		to time.Time,
	) (ContextData, error)
}
