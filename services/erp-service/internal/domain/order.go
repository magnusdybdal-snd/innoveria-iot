// Package domain defines core ERP ingest entities and business enums.
package domain

import (
	"context"
	"time"
)

// OrderStatus describes the lifecycle state of a manufacturing order.
type OrderStatus string

const (
	// OrderStatusNotInitialized means the order is created but not initialized.
	OrderStatusNotInitialized OrderStatus = "not_initialized"
	// OrderStatusRegistered means the order is registered in ERP.
	OrderStatusRegistered OrderStatus = "registered"
	// OrderStatusPrinted means order documents have been printed.
	OrderStatusPrinted OrderStatus = "printed"
	// OrderStatusStarted means work has started on the order.
	OrderStatusStarted OrderStatus = "started"
	// OrderStatusFinished means production is finished.
	OrderStatusFinished OrderStatus = "finished"
	// OrderStatusPostCalculated means post-calculation has been completed.
	OrderStatusPostCalculated OrderStatus = "post_calculated"
	// OrderStatusDelivered means the order has been delivered.
	OrderStatusDelivered OrderStatus = "delivered"
	// OrderStatusHistorical means the order is archived as historical.
	OrderStatusHistorical OrderStatus = "historical"
)

// OrderSummary represent all available orders from Monitor ERP
// its used for showcasing all available orderIDs, so the orderID can be reused for aggregated result
type OrderSummary struct {
	ID          int64
	OrderNumber string
}

// OrderAggregate represents a high-level aggregation of orders and their
// associated operations. It is typically used as a composite structure
// to group multiple orders together with all related operational data.
type OrderAggregate struct {
	Order               []Order
	Operations          []OrderOperationWithReports
	ProductionResources []ProductionResource
}

// OrderOperationWithReports represents a single operation within an order,
// enriched with its execution reports and the production resource used.
// It is useful for tracking both the planned operation and its actual
// execution details in one structure.
type OrderOperationWithReports struct {
	Operation OrderOperation
	Reports   []OrderReport
}

// Order represents a manufacturing order ingested from Monitor ERP.
type Order struct {
	ID                int64
	CompanyID         string
	OrderNumber       string
	PartID            string
	PartDescription   string // Human readable product name
	PlannedStartDate  time.Time
	PlannedFinishDate time.Time
	ActualStartDate   *time.Time
	ActualFinishDate  *time.Time
	Status            OrderStatus
	Priority          int
	ReceivedAt        time.Time
}

// OrderRepo defines the repository interface for orders.
type OrderRepo interface {
	FindAllByCompanyID(ctx context.Context, companyID string) ([]OrderSummary, error)
	FindByID(ctx context.Context, orderID int64, companyID string) (Order, error)
}

// OrderService defines the business logic interface for orders
type OrderService interface {
	GetOrderSummary(ctx context.Context, companyID string) ([]OrderSummary, error)
	GetOne(ctx context.Context, orderID int64, companyID string) (OrderAggregate, error)
}
