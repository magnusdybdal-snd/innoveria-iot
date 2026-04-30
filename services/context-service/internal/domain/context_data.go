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

// OrderContext is the aggregated result for a single ERP order, enriched with
// sensor data for each of its operations.
type OrderContext struct {
	Order      ERPOrder
	Operations []OperationContext
}

// OperationContext pairs a single manufacturing operation with the sensors
// assigned to its production resource.
type OperationContext struct {
	Operation ERPOrderOperation
	Sensors   []SensorContext
	// Degraded is true when sensor data could not be loaded due to a technical
	// error (e.g. device-service timeout or 5xx). It is false when the operation
	// simply has no sensors mapped, which is a valid expected state.
	Degraded bool
}

// SensorContext holds a sensor, its metric definitions, and its raw measurements
// within the order's actual time window.
type SensorContext struct {
	Sensor       DeviceSensor
	Metrics      []SensorMetric
	Measurements []MeasurementReading
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
	// GetOrders retrieves a slim list of ERP orders for populating the order picker.
	GetOrders(ctx context.Context, companyID, userID, role string) ([]ERPOrderSummary, error)
	// GetOrderByID retrieves a single ERP order with full detail (operations, production resources, reports).
	// Returns ErrNotFound if no order matches.
	GetOrderByID(ctx context.Context, companyID, userID, role string, orderID int64) (*ERPOrder, error)
	// GetOrderContext aggregates ERP order data with sensor readings for a single order.
	GetOrderContext(ctx context.Context, companyID, userID, role string, orderID int64) (*OrderContext, error)

	// GetProductionResources returns all ERP production resources (work centers) for the given company.
	//
	// TODO: replace with AUTH — companyID should be read from the gateway-injected X-Auth-Company-Id header once auth middleware propagation is wired up.
	GetProductionResources(ctx context.Context, companyID string) ([]ERPProductionResource, error)
}
