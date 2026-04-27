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

// ERPOrderReport is a single production reporting event linked to an operation.
type ERPOrderReport struct {
	ID                 int64
	Quantity           float64
	RestQuantity       float64
	Type               string
	ReportingTimestamp time.Time
	ActualReportedDate *time.Time
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
	Status                   string
	ProductionResourceStatus string
	Reports                  []ERPOrderReport
}

// ERPOrder represents a manufacturing order from Monitor ERP, enriched with
// its operations, production resources, and reporting events.
// ReceivedAt reflects when the erp-agent last synced this order.
type ERPOrder struct {
	ID                int64
	OrderNumber       string
	PartID            string
	PartDescription   string
	PlannedStartDate  time.Time
	PlannedFinishDate time.Time
	ActualStartDate   *time.Time
	ActualFinishDate  *time.Time
	Status            string
	Priority          int
	Operations        []ERPOrderOperation
	ReceivedAt        time.Time
}
