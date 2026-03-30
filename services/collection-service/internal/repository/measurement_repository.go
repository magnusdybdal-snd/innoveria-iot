// Package repository implements the persistence layer for the collection service.
package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"innoveria-iot/collection-service/internal/db"
	"innoveria-iot/collection-service/internal/domain"
	"time"
)

const (
	insertMeasurementQuery = `
		INSERT INTO collection.sensor_measurement (device_eui, timestamp, payload, company_id)
		SELECT $1, $2, $3, company_id
		FROM collection.tenant_mapping
		WHERE chirpstack_tenant_id = $4
	`

	findLatestQuery = `
		SELECT device_eui, timestamp, payload, company_id
		FROM collection.sensor_measurement
		WHERE device_eui = $1
		ORDER BY timestamp DESC
		LIMIT 1
	`

	findByTimeRangeQuery = `
		SELECT device_eui, timestamp, payload, company_id
		FROM collection.sensor_measurement
		WHERE device_eui = $1
		  AND timestamp >= $2
		  AND timestamp <= $3
		ORDER BY timestamp ASC
	`

	findPayloadKeysQuery = `
		SELECT DISTINCT jsonb_object_keys(payload)
		FROM (
			SELECT payload
			FROM collection.sensor_measurement
			WHERE device_eui = $1
			AND payload IS NOT NULL
			AND jsonb_typeof(payload) = 'object'
			ORDER BY timestamp DESC
			LIMIT 10
		) recent
		ORDER BY 1 ASC
	`
)

// MeasurementRepository handles persistence of sensor measurements against a TimescaleDB hypertable.
type MeasurementRepository struct {
	db *db.DB
}

// NewMeasurementRepository creates a new MeasurementRepository backed by the given database connection.
func NewMeasurementRepository(db *db.DB) *MeasurementRepository {
	return &MeasurementRepository{db: db}
}

// Insert stores a measurement, resolving company_id from the tenant mapping using the Chirpstack tenantID.
func (r *MeasurementRepository) Insert(ctx context.Context, measurement domain.SensorMeasurement, tenantID string) error {
	// Marshal the payload into JSONB (json bytes) for storage in database
	// The payload shape will vary depending on the sensor and codec in Chirpstack
	payloadJSON, err := json.Marshal(measurement.Payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	// Checks the response from postgres
	tag, err := r.db.Pool.Exec(ctx, insertMeasurementQuery, measurement.DeviceEUI, measurement.Timestamp, payloadJSON, tenantID)
	if err != nil {
		return fmt.Errorf("insert measurement %w", err)
	}

	// If no rows where inserted (no tenant mapping found for Chirpstack tenant)
	// We throw an error. This means the device is not linked to any company
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("unknown tenant: %s", tenantID)
	}

	return nil
}

// FindLatest returns the most recent measurement for the given device.
func (r *MeasurementRepository) FindLatest(ctx context.Context, deviceEUI string) (domain.SensorMeasurement, error) {

	var measurement domain.SensorMeasurement
	// pgx cannot scan JSONB directly into our map structure, so it needs to be unmarshaled first.
	var payloadBytes []byte

	err := r.db.Pool.QueryRow(ctx, findLatestQuery, deviceEUI).Scan(
		&measurement.DeviceEUI,
		&measurement.Timestamp,
		&payloadBytes,
		&measurement.CompanyID,
	)

	if err != nil {
		return domain.SensorMeasurement{}, fmt.Errorf("find latest: %w", err)
	}

	if err := json.Unmarshal(payloadBytes, &measurement.Payload); err != nil {
		return domain.SensorMeasurement{}, fmt.Errorf("unmarshal payload: %w", err)
	}

	return measurement, nil

}

// FindByTimeRange returns all measurements for a device within the given time window, ordered oldest first.
func (r *MeasurementRepository) FindByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]domain.SensorMeasurement, error) {

	rows, err := r.db.Pool.Query(ctx, findByTimeRangeQuery, deviceEUI, from, to)
	if err != nil {
		return nil, fmt.Errorf("find by time range: %w", err)
	}
	defer rows.Close()

	var measurements []domain.SensorMeasurement

	for rows.Next() {
		var measurement domain.SensorMeasurement
		// pgx cannot scan JSONB directly into our map structure, so it needs to be unmarshaled first.
		var payloadBytes []byte

		if err := rows.Scan(&measurement.DeviceEUI, &measurement.Timestamp, &payloadBytes, &measurement.CompanyID); err != nil {
			return nil, fmt.Errorf("scan measurement: %w", err)
		}

		if err := json.Unmarshal(payloadBytes, &measurement.Payload); err != nil {
			return nil, fmt.Errorf("unmarshal payload: %w", err)
		}

		measurements = append(measurements, measurement)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return measurements, nil
}

// FindPayloadKeys retrieves unique payload keys from the last 10 readings for a device based on deviceEUI
func (r *MeasurementRepository) FindPayloadKeys(ctx context.Context, deviceEUI string) ([]string, error) {
	rows, err := r.db.Pool.Query(ctx, findPayloadKeysQuery, deviceEUI)
	if err != nil {
		return nil, fmt.Errorf("find payload keys for %s: %w", deviceEUI, err)
	}
	defer rows.Close()

	out := []string{}

	for rows.Next() {
		var key string
		if err := rows.Scan(&key); err != nil {
			return nil, fmt.Errorf("find payload keys for %s: scan payload key: %w", deviceEUI, err)
		}
		out = append(out, key)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("find payload keys for %s: rows error: %w", deviceEUI, err)
	}

	return out, nil
}
