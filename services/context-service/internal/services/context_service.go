// Package services contains context-service business logic.
package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"

	"innoveria-iot/context-service/internal/calculators"
	"innoveria-iot/context-service/internal/domain"
)

// ContextServiceImpl implements context use cases for context service.
type ContextServiceImpl struct {
	collectionClient domain.CollectionClient
	erpClient        domain.ERPClient
	deviceClient     domain.DeviceClient
	ruleRepo         domain.RuleRepository
	calculators      map[string]calculators.Calculator
}

// NewContextServiceImpl creates a new ContextServiceImpl instance.
func NewContextServiceImpl(collectionClient domain.CollectionClient, erpClient domain.ERPClient, deviceClient domain.DeviceClient, ruleRepo domain.RuleRepository) *ContextServiceImpl {
	return &ContextServiceImpl{
		collectionClient: collectionClient,
		erpClient:        erpClient,
		deviceClient:     deviceClient,
		ruleRepo:         ruleRepo,
		calculators:      calculators.NewRegistry(),
	}
}

// GetContextData fetches raw measurements for each device and computes context data
// using the aggregation rule identified by ruleID.
// bucketMins overrides the rule's default TimeBucketMinutes when greater than zero.
func (s *ContextServiceImpl) GetContextData(
	ctx context.Context,
	companyID string,
	deviceEUIs []string,
	ruleID string,
	from, to time.Time,
	bucketMins int,
) ([]domain.ContextData, error) {

	// Look up and validate the rule
	rule, err := s.ruleRepo.GetByID(ctx, ruleID)
	if err != nil {
		return nil, err
	}
	if rule.CompanyID != companyID {
		// Return ErrNotFound to avoid leaking that the rule exists but belongs to another company
		return nil, domain.ErrNotFound
	}

	// Resolve effective bucket size: caller override takes precedence over rule default
	effectiveBucketMins := rule.TimeBucketMinutes
	if bucketMins > 0 {
		effectiveBucketMins = bucketMins
	}

	// Pick calculator: fall back to generic aggregation when no specific one is registered
	calc, ok := s.calculators[rule.ContextType]
	if !ok {
		slog.Warn("no calculator registered for context type, falling back to generic aggregation", "context_type", rule.ContextType, "rule_id", ruleID)
		calc = &calculators.GenericAggregationCalculator{}
	}

	results := make([]domain.ContextData, 0, len(deviceEUIs))
	for _, eui := range deviceEUIs {
		// TODO: ensure that tenant scoping is applied in the collection (when ready)
		readings, err := s.collectionClient.GetMeasurements(ctx, eui, from, to)
		if err != nil {
			return nil, fmt.Errorf("fetching measurements for device %s: %w", eui, err)
		}

		// A device with no data in the window is not an error — return a zero result
		if len(readings) == 0 {
			results = append(results, domain.ContextData{
				CompanyID:    companyID,
				DeviceEUI:    eui,
				ContextType:  rule.ContextType,
				Unit:         rule.MeasurementType,
				PeriodStart:  from,
				PeriodEnd:    to,
				Buckets:      []domain.ContextBucket{},
				CalculatedAt: time.Now().UTC(),
			})
			continue
		}

		result, err := calc.Calculate(calculators.Input{
			Rule:          rule,
			Readings:      readings,
			BucketMinutes: effectiveBucketMins,
			From:          from,
			To:            to,
		})
		if err != nil {
			return nil, fmt.Errorf("calculating context for device %s: %w", eui, err)
		}

		result.DeviceEUI = eui
		results = append(results, result)
	}

	return results, nil
}

// GetOrders retrieves all ERP orders enriched with their operations and
// production resources. The companyID is used to scope the query.
//
// TODO: replace with AUTH — use X-Auth-Company-Id header once auth middleware propagation is wired up end-to-end.
func (s *ContextServiceImpl) GetOrders(ctx context.Context, companyID string) ([]domain.ERPOrder, error) {
	orders, err := s.erpClient.GetOrders(ctx, companyID)
	if err != nil {
		return nil, fmt.Errorf("fetching orders from ERP: %w", err)
	}
	return orders, nil
}

// GetOrderContext aggregates ERP order data with sensor and measurement context
// for a single order identified by orderID.
//
// TODO: replace with AUTH — companyID is hardcoded until auth middleware propagation is wired up.
func (s *ContextServiceImpl) GetOrderContext(ctx context.Context, companyID string, orderID int64) (*domain.OrderContext, error) {
	found, err := s.erpClient.GetOrderByID(ctx, companyID, orderID)
	if err != nil {
		return nil, fmt.Errorf("fetching order from ERP: %w", err)
	}

	ops := make([]domain.OperationContext, len(found.Operations))

	// Guard: if the order has no actual time window, return operations without sensor data.
	if found.ActualStartDate == nil || found.ActualFinishDate == nil {
		for i, op := range found.Operations {
			ops[i] = domain.OperationContext{
				Operation: op,
				Sensors:   []domain.SensorContext{},
			}
		}
		return &domain.OrderContext{
			Order:           *found,
			Operations:      ops,
			HasMeasurements: false,
		}, nil
	}

	from := *found.ActualStartDate
	to := *found.ActualFinishDate

	g, gctx := errgroup.WithContext(ctx)

	for i, op := range found.Operations {
		g.Go(func() error {
			productionResourceID := strconv.FormatInt(op.ProductionResource.ID, 10)

			sensors, err := s.deviceClient.GetSensorsByProductionResourceID(gctx, productionResourceID)
			if err != nil {
				if errors.Is(err, domain.ErrNotFound) {
					// No sensors mapped to this production resource — expected, not an error.
					ops[i] = domain.OperationContext{Operation: op, Sensors: []domain.SensorContext{}, Degraded: false}
					return nil
				}
				// Technical failure (timeout, 5xx) — return partial data but mark as degraded.
				slog.Warn("failed to fetch sensors for production resource, returning degraded operation", "production_resource_id", productionResourceID, "error", err)
				ops[i] = domain.OperationContext{Operation: op, Sensors: []domain.SensorContext{}, Degraded: true}
				return nil
			}

			sensorContexts := make([]domain.SensorContext, len(sensors))
			sg, sgctx := errgroup.WithContext(gctx)

			for j, sensor := range sensors {
				sg.Go(func() error {
					metrics, err := s.deviceClient.GetSensorMetrics(sgctx, sensor.DeviceEUI)
					if err != nil {
						return fmt.Errorf("fetching metrics for sensor %s: %w", sensor.DeviceEUI, err)
					}

					measurements, err := s.collectionClient.GetMeasurements(sgctx, sensor.DeviceEUI, from, to)
					if err != nil {
						return fmt.Errorf("fetching measurements for sensor %s: %w", sensor.DeviceEUI, err)
					}

					sensorContexts[j] = domain.SensorContext{
						Sensor:       sensor,
						Metrics:      metrics,
						Measurements: measurements,
					}
					return nil
				})
			}

			if err := sg.Wait(); err != nil {
				return err
			}

			ops[i] = domain.OperationContext{
				Operation: op,
				Sensors:   sensorContexts,
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &domain.OrderContext{
		Order:           *found,
		Operations:      ops,
		HasMeasurements: true,
	}, nil
}
