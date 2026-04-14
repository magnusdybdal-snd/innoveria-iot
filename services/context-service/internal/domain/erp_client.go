package domain

import (
	"context"
	"time"
)

// ERPClient defines the interface for fetching order data from the ERP service.
// The interface is intentionally minimal — only the fields the context service
// needs are represented here.
//
// TODO: expand fields once the real ERP endpoint shapes are finalised.
type ERPClient interface {
	// GetOrders returns a summary list of all orders for the given company.
	GetOrders(ctx context.Context, companyID string) ([]ERPOrder, error)
	// GetOrderDetail returns the full detail for a single order, including
	// its reportings and workcenter information.
	GetOrderDetail(ctx context.Context, orderID string) (*ERPOrderDetail, error)
}

// ERPOrder is a summary item in the orders list returned by the ERP service.
//
// TODO: replace with AUTH — companyID scoping will be enforced by the gateway once the auth middleware propagation is fully wired up.
type ERPOrder struct {
	OrderID        string
	ProductName    string
	Status         string
	StartTime      time.Time
	EndTime        time.Time
	WorkcenterID   string
	WorkcenterName string
}

// ERPOrderReporting is a single production reporting event for an order.
type ERPOrderReporting struct {
	ReportingID string
	Timestamp   time.Time
	Quantity    float64
	Status      string
}

// ERPOrderDetail is the full detail for a single order, including its reporting
// events and workcenter reference.
type ERPOrderDetail struct {
	ERPOrder
	Reportings []ERPOrderReporting
}
