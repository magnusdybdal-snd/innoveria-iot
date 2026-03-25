package service

import (
	"context"
	"fmt"
	"log/slog"

	"innoveria-iot/device-service/internal/domain"
)

// SensorMetricServiceImpl implements domain.SensorMetricService.
type SensorMetricServiceImpl struct {
	sensorRepo        domain.SensorRepository
	sensorMetricRepo  domain.SensorMetricRepository
	payloadSchemaRepo domain.PayloadSchemaRepository
}

// NewSensorMetricService creates a new SensorMetricServiceImpl with the given repositories.
func NewSensorMetricService(
	sensorRepo domain.SensorRepository,
	sensorMetricRepo domain.SensorMetricRepository,
	payloadSchemaRepo domain.PayloadSchemaRepository,
) *SensorMetricServiceImpl {
	return &SensorMetricServiceImpl{
		sensorRepo:        sensorRepo,
		sensorMetricRepo:  sensorMetricRepo,
		payloadSchemaRepo: payloadSchemaRepo,
	}
}

// UpsertMetrics saves operator-defined metric labels for a configurable sensor.
// Returns domain.ErrInvalidMeasurementType if any slug does not exist in the vocabulary.
func (s *SensorMetricServiceImpl) UpsertMetrics(ctx context.Context, deviceEUI string, metrics []domain.SensorMetric) error {
	// Guard, validation is handled in handler, if no metrics comes in, nothign to save.
	if len(metrics) == 0 {
		return nil
	}

	sensor, err := s.sensorRepo.FindByEUI(ctx, deviceEUI)
	if err != nil {
		return err
	}

	for i := range metrics {
		metrics[i].SensorID = sensor.Id
	}

	if err := s.sensorMetricRepo.UpsertBatch(ctx, metrics); err != nil {
		return err
	}
	slog.Info("upserted sensor metrics", "sensor_id", sensor.Id, "count", len(metrics))
	return nil
}

// GetEffectiveMetrics resolves the payload key mappings for a sensor.
// Checks per-sensor metrics first (configurable sensors), falls back to the profile-level
// payload schema (fixed-schema sensors). Returns an empty slice if not yet configured.
func (s *SensorMetricServiceImpl) GetEffectiveMetrics(ctx context.Context, deviceEUI string) ([]domain.SensorMetric, error) {
	sensor, err := s.sensorRepo.FindByEUI(ctx, deviceEUI)
	if err != nil {
		return nil, fmt.Errorf("get effective metrics: find sensor: %w", err)
	}

	// Check per-sensor metrics first (configurable sensors).
	metrics, err := s.sensorMetricRepo.FindBySensorID(ctx, sensor.Id)
	if err != nil {
		return nil, fmt.Errorf("get effective metrics: find sensor metrics: %w", err)
	}
	if len(metrics) > 0 {
		return metrics, nil
	}

	// Fall back to profile-level payload schema (fixed-schema sensors).
	schemas, err := s.payloadSchemaRepo.FindByProfileID(ctx, sensor.ChirpstackProfileID)
	if err != nil {
		return nil, fmt.Errorf("get effective metrics: find payload schema: %w", err)
	}

	// Only return labeled rows — drafts (nil MeasurementType) are not actionable.
	out := []domain.SensorMetric{}
	for _, schema := range schemas {
		if schema.MeasurementType == nil {
			continue
		}
		out = append(out, domain.SensorMetric{
			PayloadKey:      schema.PayloadKey,
			MeasurementType: *schema.MeasurementType,
			Unit:            schema.Unit,
		})
	}

	return out, nil
}
