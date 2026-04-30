// Package services contains context-service business logic.
package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"sync"

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
			Rule:              rule,
			Readings:          readings,
			BucketMinutes:     effectiveBucketMins,
			From:              from,
			To:                to,
			CurrentPayloadKey: rule.MeasurementType,
			// TODO: look up per-sensor voltage when GetContextData callers need it.
			VoltageV: 230.0,
		})
		if err != nil {
			return nil, fmt.Errorf("calculating context for device %s: %w", eui, err)
		}

		result.DeviceEUI = eui
		results = append(results, result)
	}

	return results, nil
}

// GetOrders retrieves a slim list of ERP orders for populating the order picker.
func (s *ContextServiceImpl) GetOrders(ctx context.Context, companyID, userID, role string) ([]domain.ERPOrderSummary, error) {
	orders, err := s.erpClient.GetOrders(ctx, companyID, userID, role)
	if err != nil {
		return nil, fmt.Errorf("fetching orders from ERP: %w", err)
	}
	return orders, nil
}

// GetOrderByID retrieves a single ERP order with full detail.
func (s *ContextServiceImpl) GetOrderByID(ctx context.Context, companyID, userID, role string, orderID int64) (*domain.ERPOrder, error) {
	order, err := s.erpClient.GetOrderByID(ctx, companyID, userID, role, orderID)
	if err != nil {
		return nil, fmt.Errorf("fetching order from ERP: %w", err)
	}
	return order, nil
}

// GetProductionResources retrieves all ERP production resources (work centers) for the given company.
func (s *ContextServiceImpl) GetProductionResources(ctx context.Context, companyID string) ([]domain.ERPProductionResource, error) {
	resources, err := s.erpClient.GetProductionResources(ctx, companyID)
	if err != nil {
		return nil, fmt.Errorf("fetching production resources from ERP: %w", err)
	}
	return resources, nil
}

// GetOrderContext aggregates ERP order data with sensor and measurement context
// for a single order identified by orderID.
func (s *ContextServiceImpl) GetOrderContext(ctx context.Context, companyID, userID, role string, orderID int64) (*domain.OrderContext, error) {
	found, err := s.erpClient.GetOrderByID(ctx, companyID, userID, role, orderID)
	if err != nil {
		return nil, fmt.Errorf("fetching order from ERP: %w", err)
	}

	ops := make([]domain.OperationContext, len(found.Operations))

	// Guard: if the order has no actual time window, return operations without sensor data.
	if found.ActualStartDate == nil || found.ActualFinishDate == nil {
		for i, op := range found.Operations {
			ops[i] = domain.OperationContext{Operation: op, Sensors: []domain.SensorContext{}}
		}
		return &domain.OrderContext{Order: *found, Operations: ops}, nil
	}

	from := *found.ActualStartDate
	to := *found.ActualFinishDate

	var wg sync.WaitGroup
	for i, op := range found.Operations {
		wg.Go(func() {
			ops[i] = s.buildOperationContext(ctx, op, from, to)
		})
	}
	wg.Wait()

	return &domain.OrderContext{Order: *found, Operations: ops}, nil
}

// buildOperationContext fetches sensors for an operation and enriches each with
// metrics and measurements over the given time window.
func (s *ContextServiceImpl) buildOperationContext(ctx context.Context, op domain.ERPOrderOperation, from, to time.Time) domain.OperationContext {
	productionResourceID := strconv.FormatInt(op.ProductionResource.ID, 10)

	sensors, err := s.deviceClient.GetSensorsByProductionResourceID(ctx, productionResourceID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			// No sensors mapped to this production resource — expected, not an error.
			return domain.OperationContext{Operation: op, Sensors: []domain.SensorContext{}}
		}
		// Technical failure (timeout, 5xx) — return partial data but mark as degraded.
		slog.Error("failed to fetch sensors for production resource, returning degraded operation", "production_resource_id", productionResourceID, "error", err)
		return domain.OperationContext{Operation: op, Sensors: []domain.SensorContext{}, Degraded: true}
	}

	sensorContexts := make([]domain.SensorContext, len(sensors))
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		svcErr error
	)

	for j, sensor := range sensors {
		wg.Go(func() {
			sc, err := s.buildSensorContext(ctx, sensor, from, to)
			if err != nil {
				mu.Lock()
				if svcErr == nil {
					svcErr = err
				}
				mu.Unlock()
				return
			}
			sensorContexts[j] = sc
		})
	}
	wg.Wait()

	if svcErr != nil {
		slog.Error("failed to build sensor context for operation, returning degraded operation", "production_resource_id", productionResourceID, "error", svcErr)
		return domain.OperationContext{Operation: op, Sensors: []domain.SensorContext{}, Degraded: true}
	}

	return domain.OperationContext{Operation: op, Sensors: sensorContexts}
}

// buildSensorContext fetches metrics and measurements for a single sensor over the given time window.
// For electricity sensors with a configured voltage, it also computes total energy consumption in Wh.
func (s *ContextServiceImpl) buildSensorContext(ctx context.Context, sensor domain.DeviceSensor, from, to time.Time) (domain.SensorContext, error) {
	metrics, err := s.deviceClient.GetSensorMetrics(ctx, sensor.DeviceEUI)
	if err != nil {
		return domain.SensorContext{}, fmt.Errorf("fetching metrics for sensor %s: %w", sensor.DeviceEUI, err)
	}

	measurements, err := s.collectionClient.GetMeasurements(ctx, sensor.DeviceEUI, from, to)
	if err != nil {
		return domain.SensorContext{}, fmt.Errorf("fetching measurements for sensor %s: %w", sensor.DeviceEUI, err)
	}

	sc := domain.SensorContext{Sensor: sensor, Metrics: metrics, Measurements: measurements}

	if sensor.ElectricitySensor {
		if sensor.Voltage == nil {
			slog.Warn("electricity sensor has no voltage configured, skipping Wh calculation", "device_eui", sensor.DeviceEUI)
		} else {
			currentKey := findPayloadKeyByMeasurementType(metrics, "electric_current")
			if currentKey == "" {
				slog.Warn("electricity sensor has no electric_current metric, skipping Wh calculation", "device_eui", sensor.DeviceEUI)
			} else {
				calc := &calculators.WattHourCalculator{}
				result, err := calc.Calculate(calculators.Input{
					Readings:          measurements,
					From:              from,
					To:                to,
					VoltageV:          float64(*sensor.Voltage),
					CurrentPayloadKey: currentKey,
				})
				if err != nil {
					slog.Error("Wh calculation failed for sensor", "device_eui", sensor.DeviceEUI, "error", err)
				} else {
					sc.PowerConsumptionWh = &result.Value
				}
			}
		}
	}

	return sc, nil
}

// findPayloadKeyByMeasurementType returns the payload key for the first metric matching the given
// measurement type slug, or an empty string if none is found.
func findPayloadKeyByMeasurementType(metrics []domain.SensorMetric, measurementType string) string {
	for _, m := range metrics {
		if m.MeasurementType == measurementType {
			return m.PayloadKey
		}
	}
	return ""
}
