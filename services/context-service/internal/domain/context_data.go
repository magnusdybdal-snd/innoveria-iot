package domain

import (
	"context"
	"time"
)

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
	// GetOrders retrieves all ERP orders enriched with their full detail.
	//
	// TODO: replace with AUTH — companyID should be read from the gateway-injected X-Auth-Company-Id header once auth middleware propagation is wired up.
	GetOrders(ctx context.Context, companyID string) ([]ERPOrderDetail, error)
}
