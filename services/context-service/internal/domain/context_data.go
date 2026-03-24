package domain

import (
	"context"
	"errors"
	"time"
)

// ErrNotFound is returned when a requested resource does not exist.
var ErrNotFound = errors.New("not found")

// ErrConflict is returned when a resource already exists and cannot be duplicated.
var ErrConflict = errors.New("conflict")

// ContextBucket represents a single time bucket in a context data result.
type ContextBucket struct {
	PeriodStart time.Time
	PeriodEnd   time.Time
	Value       float64
}

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
	Buckets              []ContextBucket
	CalculatedAt         time.Time
}

// ContextService defines the context service interface for the context domain.
type ContextService interface {
	GetContextData(
		ctx context.Context,
		companyID string,
		deviceEUIs []string,
		ruleID string,
		from, to time.Time,
		bucketMins int,
	) ([]ContextData, error)
}
