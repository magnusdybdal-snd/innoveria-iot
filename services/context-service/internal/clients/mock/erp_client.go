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

var (
	desc1 = "Welding Station 1"
	desc2 = "Press Station 2"
	desc3 = "Assembly Line A"
)

var mockOrders = []domain.ERPOrder{
	{
		ID:                1,
		OrderNumber:       "MO-2026-001",
		PartDescription:   "Steel Frame A",
		PlannedStartDate:  time.Date(2026, 4, 14, 6, 0, 0, 0, time.UTC),
		PlannedFinishDate: time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
		ActualStartDate:   ptr(time.Date(2026, 4, 14, 6, 12, 0, 0, time.UTC)),
		ActualFinishDate:  nil,
		Status:            "started",
		Priority:          1,
		Operations: []domain.ERPOrderOperation{
			{
				ID: 10,
				ProductionResource: domain.ERPProductionResource{
					ID:          1,
					Number:      "WC-101",
					Description: &desc1,
					Type:        "machine",
				},
				PlannedStartDate:         time.Date(2026, 4, 14, 6, 0, 0, 0, time.UTC),
				PlannedFinishDate:        time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
				ActualStartDate:          ptr(time.Date(2026, 4, 14, 6, 12, 0, 0, time.UTC)),
				ActualFinishDate:         nil,
				Status:                   "started",
				ProductionResourceStatus: "started",
			},
		},
	},
	{
		ID:                2,
		OrderNumber:       "MO-2026-002",
		PartDescription:   "Aluminium Panel B",
		PlannedStartDate:  time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
		PlannedFinishDate: time.Date(2026, 4, 14, 22, 0, 0, 0, time.UTC),
		ActualStartDate:   nil,
		ActualFinishDate:  nil,
		Status:            "registered",
		Priority:          2,
		Operations: []domain.ERPOrderOperation{
			{
				ID: 20,
				ProductionResource: domain.ERPProductionResource{
					ID:          2,
					Number:      "WC-102",
					Description: &desc2,
					Type:        "machine",
				},
				PlannedStartDate:         time.Date(2026, 4, 14, 14, 0, 0, 0, time.UTC),
				PlannedFinishDate:        time.Date(2026, 4, 14, 22, 0, 0, 0, time.UTC),
				ActualStartDate:          nil,
				ActualFinishDate:         nil,
				Status:                   "none",
				ProductionResourceStatus: "none",
			},
		},
	},
	{
		ID:                3,
		OrderNumber:       "MO-2026-003",
		PartDescription:   "Copper Coil C",
		PlannedStartDate:  time.Date(2026, 4, 13, 6, 0, 0, 0, time.UTC),
		PlannedFinishDate: time.Date(2026, 4, 13, 14, 0, 0, 0, time.UTC),
		ActualStartDate:   ptr(time.Date(2026, 4, 13, 6, 5, 0, 0, time.UTC)),
		ActualFinishDate:  ptr(time.Date(2026, 4, 13, 13, 50, 0, 0, time.UTC)),
		Status:            "finished",
		Priority:          1,
		Operations: []domain.ERPOrderOperation{
			{
				ID: 30,
				ProductionResource: domain.ERPProductionResource{
					ID:          1,
					Number:      "WC-101",
					Description: &desc1,
					Type:        "machine",
				},
				PlannedStartDate:         time.Date(2026, 4, 13, 6, 0, 0, 0, time.UTC),
				PlannedFinishDate:        time.Date(2026, 4, 13, 10, 0, 0, 0, time.UTC),
				ActualStartDate:          ptr(time.Date(2026, 4, 13, 6, 5, 0, 0, time.UTC)),
				ActualFinishDate:         ptr(time.Date(2026, 4, 13, 9, 55, 0, 0, time.UTC)),
				Status:                   "finished",
				ProductionResourceStatus: "finished",
			},
			{
				ID: 31,
				ProductionResource: domain.ERPProductionResource{
					ID:          3,
					Number:      "WC-103",
					Description: &desc3,
					Type:        "manual_work",
				},
				PlannedStartDate:         time.Date(2026, 4, 13, 10, 0, 0, 0, time.UTC),
				PlannedFinishDate:        time.Date(2026, 4, 13, 14, 0, 0, 0, time.UTC),
				ActualStartDate:          ptr(time.Date(2026, 4, 13, 10, 10, 0, 0, time.UTC)),
				ActualFinishDate:         ptr(time.Date(2026, 4, 13, 13, 50, 0, 0, time.UTC)),
				Status:                   "finished",
				ProductionResourceStatus: "finished",
			},
			{
				// Edge case: numeric-looking status (2-digit number)
				ID: 32,
				ProductionResource: domain.ERPProductionResource{
					ID:          4,
					Number:      "WC-104",
					Description: nil,
					Type:        "machine",
				},
				PlannedStartDate:         time.Date(2026, 4, 13, 14, 0, 0, 0, time.UTC),
				PlannedFinishDate:        time.Date(2026, 4, 13, 16, 0, 0, 0, time.UTC),
				ActualStartDate:          ptr(time.Date(2026, 4, 13, 14, 5, 0, 0, time.UTC)),
				ActualFinishDate:         ptr(time.Date(2026, 4, 13, 15, 58, 0, 0, time.UTC)),
				Status:                   "finished",
				ProductionResourceStatus: "42",
			},
			{
				// Edge case: long string status with a space
				ID: 33,
				ProductionResource: domain.ERPProductionResource{
					ID:          5,
					Number:      "WC-105",
					Description: nil,
					Type:        "machine",
				},
				PlannedStartDate:         time.Date(2026, 4, 13, 16, 0, 0, 0, time.UTC),
				PlannedFinishDate:        time.Date(2026, 4, 13, 18, 0, 0, 0, time.UTC),
				ActualStartDate:          ptr(time.Date(2026, 4, 13, 16, 10, 0, 0, time.UTC)),
				ActualFinishDate:         ptr(time.Date(2026, 4, 13, 17, 55, 0, 0, time.UTC)),
				Status:                   "finished",
				ProductionResourceStatus: "thisisalongstringfortesting andthisisanewline",
			},
		},
	},
}

// GetOrders returns the mock order list. companyID is accepted but not used.
func (c *ERPClient) GetOrders(_ context.Context, _ string) ([]domain.ERPOrder, error) {
	return mockOrders, nil
}

func ptr(t time.Time) *time.Time {
	return &t
}
