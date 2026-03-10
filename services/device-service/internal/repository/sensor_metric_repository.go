package repository

import (
	"context"
	"fmt"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

const (
	upsertSensorMetricQuery = `
		INSERT INTO device.sensor_metric (sensor_id, measurement_type, unit)
		VALUES ($1, $2, $3)
		ON CONFLICT (sensor_id, measurement_type)
		DO UPDATE SET unit = EXCLUDED.unit
	`

	findSensorMetricBySensorIDQuery = `
		SELECT sensor_id, measurement_type, unit
		FROM device.sensor_metric
		WHERE sensor_id = $1
		ORDER BY measurement_type ASC
	`
)

// SensorMetricRepository handles persistence of sensor metrics stored in the database.
type SensorMetricRepository struct {
	db *dbutil.DB
}

// NewSensorMetricRepository creates a new SensorMetricRepository with the given database.
func NewSensorMetricRepository(db *dbutil.DB) *SensorMetricRepository {
	return &SensorMetricRepository{db: db}
}

// Upsert inserts a new sensor metric for a sensor. Some sensors have more than one metric,
// e.g. temperature (°C) and humidity (%RH).
// If the combination of sensorID and measurement type already exists, it updates the unit instead
// to avoid duplicate measurement types.
func (r *SensorMetricRepository) Upsert(ctx context.Context, metric domain.SensorMetric) error {

	_, err := r.db.Pool.Exec(ctx, upsertSensorMetricQuery, metric.SensorID, metric.MeasurementType, metric.Unit)

	if err != nil {
		return fmt.Errorf("upsert sensor metric: %w", err)
	}

	return nil

}

// FindBySensorID retrieves all sensor metrics for a given sensor by its sensorID, ordered by measurement type.
func (r *SensorMetricRepository) FindBySensorID(ctx context.Context, sensorID string) ([]domain.SensorMetric, error) {

	// Query the database to collect all rows
	rows, err := r.db.Pool.Query(ctx, findSensorMetricBySensorIDQuery, sensorID)
	if err != nil {
		return nil, fmt.Errorf("find sensor metrics by sensor id: %w", err)
	}
	defer rows.Close()

	// The slice of sensor metrics to be returned
	var out []domain.SensorMetric

	for rows.Next() {
		var sensorMetric domain.SensorMetric

		err := rows.Scan(
			&sensorMetric.SensorID,
			&sensorMetric.MeasurementType,
			&sensorMetric.Unit,
		)
		if err != nil {
			return nil, fmt.Errorf("scan sensor metric: %w", err)
		}

		// Add the sensor to the slice
		out = append(out, sensorMetric)
	}

	// Sanity check if the loop ended due to an error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return out, nil
}
