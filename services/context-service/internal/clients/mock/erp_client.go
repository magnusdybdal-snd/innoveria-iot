// Package mock provides development-only mock implementations of external service clients.
package mock

import (
	"context"
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

// mockWindowStart/End are evaluated once at service startup.
// The window covers the 30 days before startup, capturing all measurements
// already in the collection DB. This models a completed order.
var (
	mockWindowStart = time.Now().Add(-30 * 24 * time.Hour)
	mockWindowEnd   = time.Now()
)

var mockOrders = []domain.ERPOrder{
	{
		ID:                1,
		OrderNumber:       "MO-2026-001",
		PartDescription:   "Steel Frame A",
		PlannedStartDate:  time.Date(2026, 4, 14, 6, 0, 0, 0, time.UTC),
		PlannedFinishDate: time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
		ActualStartDate:   ptr(time.Date(2026, 4, 14, 6, 12, 0, 0, time.UTC)),
		ActualFinishDate:  ptr(time.Date(2026, 4, 14, 13, 55, 0, 0, time.UTC)),
		Status:            "finished",
		Priority:          1,
		Operations: []domain.ERPOrderOperation{
			{
				ID: 10,
				ProductionResource: domain.ERPProductionResource{
					ID:          1,
					Number:      "WC-101",
					Description: "Welding Station 1",
					Type:        "machine",
				},
				PlannedStartDate:         time.Date(2026, 4, 14, 6, 0, 0, 0, time.UTC),
				PlannedFinishDate:        time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
				ActualStartDate:          ptr(time.Date(2026, 4, 14, 6, 12, 0, 0, time.UTC)),
				ActualFinishDate:         ptr(time.Date(2026, 4, 14, 13, 55, 0, 0, time.UTC)),
				Status:                   "finished",
				ProductionResourceStatus: "finished",
			},
		},
	},
	{
		ID:                2,
		OrderNumber:       "MO-2026-002",
		PartDescription:   "Aluminium Panel B",
		PlannedStartDate:  time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
		PlannedFinishDate: time.Date(2026, 4, 14, 22, 0, 0, 0, time.UTC),
		ActualStartDate:   ptr(time.Date(2026, 4, 14, 14, 5, 0, 0, time.UTC)),
		ActualFinishDate:  ptr(time.Date(2026, 4, 14, 21, 50, 0, 0, time.UTC)),
		Status:            "finished",
		Priority:          2,
		Operations: []domain.ERPOrderOperation{
			{
				ID: 20,
				ProductionResource: domain.ERPProductionResource{
					ID:          2,
					Number:      "WC-102",
					Description: "Press Station 2",
					Type:        "machine",
				},
				PlannedStartDate:         time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
				PlannedFinishDate:        time.Date(2026, 4, 14, 22, 0, 0, 0, time.UTC),
				ActualStartDate:          ptr(time.Date(2026, 4, 14, 14, 5, 0, 0, time.UTC)),
				ActualFinishDate:         ptr(time.Date(2026, 4, 14, 21, 50, 0, 0, time.UTC)),
				Status:                   "finished",
				ProductionResourceStatus: "finished",
			},
		},
	},
	// MO-2026-003: time window covers the 30 days before service startup so any
	// recently collected measurements are included. Production resource ID 1
	// matches sensors 1 & 2 in the device-service seed (device EUIs b000000000000001
	// and b000000000000002). Resource ID 6 has no mapped sensors → demos the degraded state.
	{
		ID:                3,
		OrderNumber:       "MO-2026-003",
		PartDescription:   "Live Sensor Demo",
		PlannedStartDate:  mockWindowStart,
		PlannedFinishDate: mockWindowEnd,
		ActualStartDate:   ptr(mockWindowStart),
		ActualFinishDate:  ptr(mockWindowEnd),
		Status:            "finished",
		Priority:          1,
		Operations: []domain.ERPOrderOperation{
			{
				ID: 30,
				ProductionResource: domain.ERPProductionResource{
					ID:          1,
					Number:      "WC-101",
					Description: "Welding Station 1",
					Type:        "machine",
				},
				PlannedStartDate:         mockWindowStart,
				PlannedFinishDate:        mockWindowEnd,
				ActualStartDate:          ptr(mockWindowStart),
				ActualFinishDate:         ptr(mockWindowEnd),
				Status:                   "finished",
				ProductionResourceStatus: "finished",
			},
			{
				ID: 31,
				ProductionResource: domain.ERPProductionResource{
					ID:          6,
					Number:      "WC-106",
					Description: "Cutting Station 3",
					Type:        "machine",
				},
				PlannedStartDate:         mockWindowStart,
				PlannedFinishDate:        mockWindowEnd,
				ActualStartDate:          ptr(mockWindowStart),
				ActualFinishDate:         ptr(mockWindowEnd),
				Status:                   "finished",
				ProductionResourceStatus: "finished",
			},
		},
	},
}

// GetOrders returns a slim summary list derived from the mock orders.
func (c *ERPClient) GetOrders(_ context.Context, _ string) ([]domain.ERPOrderSummary, error) {
	summaries := make([]domain.ERPOrderSummary, len(mockOrders))
	for i, o := range mockOrders {
		summaries[i] = domain.ERPOrderSummary{ID: o.ID, Name: o.OrderNumber}
	}
	return summaries, nil
}

// GetOrderByID returns a single mock order by ID. companyID is accepted but not used.
func (c *ERPClient) GetOrderByID(_ context.Context, _ string, orderID int64) (*domain.ERPOrder, error) {
	for i := range mockOrders {
		if mockOrders[i].ID == orderID {
			return &mockOrders[i], nil
		}
	}
	return nil, domain.ErrNotFound
}

func ptr(t time.Time) *time.Time {
	return &t
}
