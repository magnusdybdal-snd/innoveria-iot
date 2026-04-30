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
		PartID:            "PART-001",
		PartDescription:   "Steel Frame A",
		PlannedStartDate:  time.Date(2026, 4, 14, 6, 0, 0, 0, time.UTC),
		PlannedFinishDate: time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
		ActualStartDate:   ptr(time.Date(2026, 4, 14, 6, 12, 0, 0, time.UTC)),
		ActualFinishDate:  ptr(time.Date(2026, 4, 14, 13, 55, 0, 0, time.UTC)),
		Status:            "finished",
		Priority:          1,
		ReceivedAt:        time.Date(2026, 4, 14, 5, 0, 0, 0, time.UTC),
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
				Reports: []domain.ERPOrderReport{
					{
						ID:                 100,
						Quantity:           8.0,
						RestQuantity:       0.0,
						Type:               "regular",
						ReportingTimestamp: time.Date(2026, 4, 14, 13, 55, 0, 0, time.UTC),
						ActualReportedDate: ptr(time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC)),
					},
				},
			},
		},
	},
	{
		ID:                2,
		OrderNumber:       "MO-2026-002",
		PartID:            "PART-002",
		PartDescription:   "Aluminium Panel B",
		PlannedStartDate:  time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
		PlannedFinishDate: time.Date(2026, 4, 14, 22, 0, 0, 0, time.UTC),
		ActualStartDate:   ptr(time.Date(2026, 4, 14, 14, 5, 0, 0, time.UTC)),
		ActualFinishDate:  ptr(time.Date(2026, 4, 14, 21, 50, 0, 0, time.UTC)),
		Status:            "finished",
		Priority:          2,
		ReceivedAt:        time.Date(2026, 4, 14, 13, 0, 0, 0, time.UTC),
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
				Reports: []domain.ERPOrderReport{
					{
						ID:                 200,
						Quantity:           12.5,
						RestQuantity:       0.0,
						Type:               "regular",
						ReportingTimestamp: time.Date(2026, 4, 14, 21, 50, 0, 0, time.UTC),
						ActualReportedDate: ptr(time.Date(2026, 4, 14, 22, 0, 0, 0, time.UTC)),
					},
				},
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
		PartID:            "PART-003",
		PartDescription:   "Live Sensor Demo",
		PlannedStartDate:  mockWindowStart,
		PlannedFinishDate: mockWindowEnd,
		ActualStartDate:   ptr(mockWindowStart),
		ActualFinishDate:  ptr(mockWindowEnd),
		Status:            "finished",
		Priority:          1,
		ReceivedAt:        mockWindowStart,
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
				Reports: []domain.ERPOrderReport{
					{
						ID:                 300,
						Quantity:           5.0,
						RestQuantity:       10.0,
						Type:               "regular",
						ReportingTimestamp: mockWindowStart.Add(8 * 24 * time.Hour),
						ActualReportedDate: ptr(mockWindowStart.Add(8*24*time.Hour + 5*time.Minute)),
					},
					{
						ID:                 301,
						Quantity:           7.0,
						RestQuantity:       3.0,
						Type:               "regular",
						ReportingTimestamp: mockWindowStart.Add(20 * 24 * time.Hour),
						ActualReportedDate: ptr(mockWindowStart.Add(20*24*time.Hour + 3*time.Minute)),
					},
				},
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
				Reports: []domain.ERPOrderReport{
					{
						ID:                 310,
						Quantity:           4.0,
						RestQuantity:       3.0,
						Type:               "regular",
						ReportingTimestamp: mockWindowStart.Add(22 * 24 * time.Hour),
						ActualReportedDate: ptr(mockWindowStart.Add(22*24*time.Hour + 2*time.Minute)),
					},
				},
			},
		},
	},
}

// GetOrders returns a slim summary list derived from the mock orders.
func (c *ERPClient) GetOrders(_ context.Context, _, _, _ string) ([]domain.ERPOrderSummary, error) {
	summaries := make([]domain.ERPOrderSummary, len(mockOrders))
	for i, o := range mockOrders {
		summaries[i] = domain.ERPOrderSummary{ID: o.ID, Name: o.OrderNumber}
	}
	return summaries, nil
}

// GetOrderByID returns a single mock order by ID. Auth params are accepted but not used.
func (c *ERPClient) GetOrderByID(_ context.Context, _, _, _ string, orderID int64) (*domain.ERPOrder, error) {
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
