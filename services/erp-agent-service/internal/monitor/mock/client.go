// Package mock contains lightweight Monitor client test doubles.
package mock

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"time"

	"innoveria-iot/erp-agent-service/internal/monitor"
	"innoveria-iot/pkg/monitor/dto"
)

// Client returns deterministic changing Monitor data for local development.
type Client struct {
	mu      sync.Mutex
	counter int64
	cycle   int64
}

// New returns an in-memory Monitor client used only for development/testing.
//
// It intentionally avoids session/auth/network behavior from the real client,
// because this mock exists to keep the poll -> ingest flow running when the
// external Monitor ERP dependency is unavailable.
func New() *Client {
	return &Client{
		cycle: 1,
	}
}

// Query mimics Monitor endpoint responses with small deterministic payloads.
//
// The returned DTOs deliberately include only fields needed by current ingest
// use-cases; many upstream Monitor fields are left out on purpose so the mock
// stays stable, readable, and cheap to maintain.
//
// opts is currently ignored because dev-mode polling does not need server-side
// filtering/paging behavior to validate the integration contract.
func (c *Client) Query(_ context.Context, path monitor.MonitorERPEndpoint, _ url.Values, out any) error {
	c.mu.Lock()
	if path == monitor.Order {
		c.counter++
		c.cycle = c.counter
	}
	tick := c.cycle
	c.mu.Unlock()

	now := time.Now().UTC()

	switch path {
	case monitor.Order:
		orders, ok := out.(*[]dto.ManufacturingOrder)
		if !ok {
			return fmt.Errorf("mock monitor expected *[]dto.ManufacturingOrder for %s", path)
		}

		plannedStart := now.Add(-2 * time.Hour)
		plannedFinish := now.Add(6 * time.Hour)
		actualStart := now.Add(-90 * time.Minute)

		*orders = []dto.ManufacturingOrder{
			{
				ID:                300 + int64(tick%8),
				OrderNumber:       fmt.Sprintf("MO-%05d", 30_000+tick),
				PartID:            fmt.Sprintf("PART-%03d", 100+(tick%20)),
				PartDescription:   "Mock production part",
				PlannedQuantity:   float64(100 + (tick % 50)),
				ReportedQuantity:  float64((tick % 40) + 10),
				RestQuantity:      float64(90 - (tick % 30)),
				PlannedStartDate:  plannedStart,
				PlannedFinishDate: plannedFinish,
				ActualStartDate:   &actualStart,
				ActualFinishDate:  nil,
				Status:            int(tick % 4),
				Priority:          int((tick % 3) + 1),
			},
		}

	case monitor.OrderReportings:
		reportings, ok := out.(*[]dto.ManufacturingOrderOperationReporting)
		if !ok {
			return fmt.Errorf("mock monitor expected *[]dto.ManufacturingOrderOperationReporting for %s", path)
		}

		*reportings = []dto.ManufacturingOrderOperationReporting{
			{
				ID:                 10_000 + tick,
				OperationId:        2_000 + tick,
				Quantity:           float64((tick%5)+1) * 2,
				WorkCenterId:       700 + (tick % 3),
				RestQuantity:       float64(50 - (tick % 10)),
				Type:               int(tick % 3),
				PreviousNodeStatus: int((tick + 1) % 4),
				ReportingTimestamp: now,
			},
		}

	case monitor.OrderOperations:
		operations, ok := out.(*[]dto.ManufacturingOrderOperation)
		if !ok {
			return fmt.Errorf("mock monitor expected *[]dto.ManufacturingOrderOperation for %s", path)
		}

		plannedStart := now.Add(-30 * time.Minute)
		plannedFinish := now.Add(2 * time.Hour)

		*operations = []dto.ManufacturingOrderOperation{
			{
				ID:                      2_000 + tick,
				WorkCenterId:            700 + (tick % 3),
				ReportedQuantity:        float64((tick % 12) + 5),
				RestQuantity:            float64(100 - (tick % 30)),
				ManufacturingOrderId:    300 + (tick % 8),
				OperationNumber:         int((tick % 6) + 1),
				PlannedStartDate:        plannedStart,
				PlannedFinishDate:       plannedFinish,
				ActualStartDate:         &plannedStart,
				Status:                  int(tick % 4),
				WorkshopOperationStatus: int((tick + 1) % 4),
			},
		}

	case monitor.Workcenters:
		workcenters, ok := out.(*[]dto.WorkCenter)
		if !ok {
			return fmt.Errorf("mock monitor expected *[]dto.WorkCenter for %s", path)
		}

		*workcenters = []dto.WorkCenter{
			{
				ID:                   700,
				Number:               "WC-700",
				Description:          "Mock Cutter",
				OperationDescription: "Cutting operations",
				Type:                 1,
			},
			{
				ID:                   701,
				Number:               "WC-701",
				Description:          "Mock Assembly",
				OperationDescription: "Assembly operations",
				Type:                 2,
			},
			{
				ID:                   702,
				Number:               "WC-702",
				Description:          "Mock QA",
				OperationDescription: "Quality checks",
				Type:                 3,
			},
		}

	default:
		return fmt.Errorf("mock monitor does not support endpoint: %s", path)
	}

	return nil
}
