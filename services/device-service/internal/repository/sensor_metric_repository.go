package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	upsertSensorMetricQuery = `
		INSERT INTO device.sensor_metric (sensor_id, payload_key, measurement_type, unit)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (sensor_id, payload_key)
		DO UPDATE SET
			measurement_type = EXCLUDED.measurement_type,
			unit             = EXCLUDED.unit
	`

	findSensorMetricBySensorIDQuery = `
		SELECT id, sensor_id, payload_key, measurement_type, unit
		FROM device.sensor_metric
		WHERE sensor_id = $1
		ORDER BY payload_key ASC
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

// Upsert inserts a sensor metric or updates measurement_type and unit if a row with the
// same (sensor_id, payload_key) already exists. This handles operator reconfiguration cleanly.
// Returns domain.ErrInvalidMeasurementType if the measurement type slug does not exist in the vocabulary.
func (r *SensorMetricRepository) Upsert(ctx context.Context, metric domain.SensorMetric) error {
	_, err := r.db.Pool.Exec(ctx, upsertSensorMetricQuery,
		metric.SensorID,
		metric.PayloadKey,
		metric.MeasurementType,
		metric.Unit,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			slog.Warn("foreign key violation on sensor metric upsert", "constraint", pgErr.ConstraintName, "detail", pgErr.Detail)
			return domain.ErrInvalidMeasurementType
		}
		return fmt.Errorf("upsert sensor metric: %w", err)
	}
	return nil
}

// FindBySensorID retrieves all sensor metrics for a given sensor ordered by payload key.
func (r *SensorMetricRepository) FindBySensorID(ctx context.Context, sensorID string) ([]domain.SensorMetric, error) {
	rows, err := r.db.Pool.Query(ctx, findSensorMetricBySensorIDQuery, sensorID)
	if err != nil {
		return nil, fmt.Errorf("find sensor metrics by sensor id: %w", err)
	}
	defer rows.Close()

	out := []domain.SensorMetric{}

	for rows.Next() {
		var m domain.SensorMetric
		if err := rows.Scan(&m.ID, &m.SensorID, &m.PayloadKey, &m.MeasurementType, &m.Unit); err != nil {
			return nil, fmt.Errorf("scan sensor metric: %w", err)
		}
		out = append(out, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return out, nil
}
