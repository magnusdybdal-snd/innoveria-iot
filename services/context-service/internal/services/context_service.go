// Package services contains context-service business logic.
package services

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"innoveria-iot/context-service/internal/calculators"
	"innoveria-iot/context-service/internal/domain"
)

// ContextServiceImpl implements context use cases for context service.
type ContextServiceImpl struct {
	collectionClient domain.CollectionClient
	ruleRepo         domain.RuleRepository
	calculators      map[string]calculators.Calculator
}

// NewContextServiceImpl creates a new ContextServiceImpl instance.
func NewContextServiceImpl(collectionClient domain.CollectionClient, ruleRepo domain.RuleRepository) *ContextServiceImpl {
	return &ContextServiceImpl{
		collectionClient: collectionClient,
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
