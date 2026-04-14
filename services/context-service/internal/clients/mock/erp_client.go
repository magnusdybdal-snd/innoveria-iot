// Package mock provides development-only mock implementations of external service clients.
package mock

import (
	"context"
	"fmt"
	"time"

	"innoveria-iot/context-service/internal/domain"
)

// ERPClient is a mock implementation of domain.ERPClient that returns
// hardcoded plausible data. It is intended for local development only,
// controlled by the USE_MOCK_ERP environment variable.
type ERPClient struct{}

// NewERPClient creates a new mock ERPClient.
func NewERPClient() *ERPClient {
	return &ERPClient{}
}

var mockOrders = []domain.ERPOrder{
	{
		OrderID:        "order-001",
		ProductName:    "Steel Frame A",
		Status:         "IN_PROGRESS",
		StartTime:      time.Date(2026, 4, 14, 6, 0, 0, 0, time.UTC),
		EndTime:        time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
		WorkcenterID:   "wc-101",
		WorkcenterName: "Welding Station 1",
	},
	{
		OrderID:        "order-002",
		ProductName:    "Aluminium Panel B",
		Status:         "PLANNED",
		StartTime:      time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
		EndTime:        time.Date(2026, 4, 14, 22, 0, 0, 0, time.UTC),
		WorkcenterID:   "wc-102",
		WorkcenterName: "Press Station 2",
	},
	{
		OrderID:        "order-003",
		ProductName:    "Copper Coil C",
		Status:         "COMPLETED",
		StartTime:      time.Date(2026, 4, 13, 6, 0, 0, 0, time.UTC),
		EndTime:        time.Date(2026, 4, 13, 14, 0, 0, 0, time.UTC),
		WorkcenterID:   "wc-101",
		WorkcenterName: "Welding Station 1",
	},
}

var mockReportings = map[string][]domain.ERPOrderReporting{
	"order-001": {
		{ReportingID: "rep-001-1", Timestamp: time.Date(2026, 4, 14, 7, 30, 0, 0, time.UTC), Quantity: 10, Status: "REPORTED"},
		{ReportingID: "rep-001-2", Timestamp: time.Date(2026, 4, 14, 10, 0, 0, 0, time.UTC), Quantity: 25, Status: "REPORTED"},
	},
	"order-002": {},
	"order-003": {
		{ReportingID: "rep-003-1", Timestamp: time.Date(2026, 4, 13, 8, 0, 0, 0, time.UTC), Quantity: 50, Status: "REPORTED"},
		{ReportingID: "rep-003-2", Timestamp: time.Date(2026, 4, 13, 12, 0, 0, 0, time.UTC), Quantity: 50, Status: "REPORTED"},
	},
}

// GetOrders returns the mock order list. companyID is accepted but not used.
func (c *ERPClient) GetOrders(_ context.Context, _ string) ([]domain.ERPOrder, error) {
	return mockOrders, nil
}

// GetOrderDetail returns mock order detail for a known order ID.
// Returns an error for unknown IDs to simulate real service behaviour.
func (c *ERPClient) GetOrderDetail(_ context.Context, orderID string) (*domain.ERPOrderDetail, error) {
	for _, o := range mockOrders {
		if o.OrderID == orderID {
			return &domain.ERPOrderDetail{
				ERPOrder:   o,
				Reportings: mockReportings[orderID],
			}, nil
		}
	}
	return nil, fmt.Errorf("mock ERP: order %q not found", orderID)
}
