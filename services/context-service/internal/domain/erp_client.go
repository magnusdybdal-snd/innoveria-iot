package domain

import (
	"context"
	"time"
)

// ERPOrderSummary is the slim representation returned by the erp-service list
// endpoint. It carries only enough data to populate an order picker in the UI.
type ERPOrderSummary struct {
	ID   int64
	Name string
}

// ERPClient defines the interface for fetching order data from the ERP service.
type ERPClient interface {
	// GetOrders returns a slim list of orders for the given company, suitable
	// for populating a dropdown or order list in the UI.
	GetOrders(ctx context.Context, companyID string) ([]ERPOrderSummary, error)

	// GetOrderByID returns a single order by its ID, enriched with operations
	// and production resources. Returns ErrNotFound if no order matches.
	//
	// TODO: erp-service GET /orders/{id} endpoint shape is not yet finalised.
	GetOrderByID(ctx context.Context, companyID string, orderID int64) (*ERPOrder, error)
}

// ERPProductionResource represents a work center from Monitor ERP.
type ERPProductionResource struct {
	ID          int64
	Number      string // human-readable machine name/identifier from Monitor ERP (e.g. "WC-101")
	Description string
	Type        string // "machine", "manual_work", "sub_contract", "pool", "pick"
}

// ERPOrderOperation is a single manufacturing operation linking an order to a
// production resource. An order may have multiple operations on different work
// centers.
type ERPOrderOperation struct {
	ID                       int64
	ProductionResource       ERPProductionResource
	PlannedStartDate         time.Time
	PlannedFinishDate        time.Time
	ActualStartDate          *time.Time
	ActualFinishDate         *time.Time
	Status                   string // plan-level status of this operation
	ProductionResourceStatus string // current activity status of the work center
}

// ERPOrder represents a manufacturing order from Monitor ERP, enriched with
// its operations and the production resources (work centers) they reference.
//
// TODO: replace with AUTH — companyID scoping will be enforced by the gateway once the auth middleware propagation is fully wired up.
type ERPOrder struct {
	ID                int64
	OrderNumber       string // human-readable order identifier (e.g. "MO-2026-001")
	PartDescription   string // human-readable product name
	PlannedStartDate  time.Time
	PlannedFinishDate time.Time
	ActualStartDate   *time.Time
	ActualFinishDate  *time.Time
	Status            string // see erp-service domain.OrderStatus for enum values
	Priority          int
	Operations        []ERPOrderOperation
}
