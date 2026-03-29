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
	FindPayloadKeysQuery = `
		SELECT DISTINCT jsonb_object_keys(payload)
		FROM (
			SELECT payload
			FROM collection.sensor_measurement
			WHERE device_eui = $1
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
func (s *MeasurementRepository) Insert(ctx context.Context, measurement domain.SensorMeasurement, tenantID string) error {
	// Marshal the payload into JSONB (json bytes) for storage in database
	// The payload shape will vary depending on the sensor and codec in Chirpstack
	payloadJSON, err := json.Marshal(measurement.Payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	// Query to insert the measurement. Resolving the internal company_id from
	// the tenant_mapping table using the Chirpstack tenantID. This ensures the
	// data is tagged with our internal company ID
	// If no match is found, no data is inserted into the table
	const QUERY = `
			INSERT INTO collection.sensor_measurement (device_eui, timestamp, payload, company_id)
			SELECT $1, $2, $3, company_id
			FROM collection.tenant_mapping
			WHERE chirpstack_tenant_id = $4
	`

	// Checks the response from postgres
	tag, err := s.db.Pool.Exec(ctx, QUERY, measurement.DeviceEUI, measurement.Timestamp, payloadJSON, tenantID)
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
func (s *MeasurementRepository) FindLatest(ctx context.Context, deviceEUI string) (domain.SensorMeasurement, error) {

	// Fetch the single most recent measurement for the given device.
	const QUERY = `
			SELECT device_eui, timestamp, payload, company_id
			FROM collection.sensor_measurement
			WHERE device_eui = $1
			ORDER BY timestamp DESC
			LIMIT 1
	`

	// The sensor measurement object to be returned
	var measurement domain.SensorMeasurement
	// JSONB object coming from postgres- Holds the raw JSONB bytes from postgres
	// pgx cannot scan JSONB directly into our map structure, so it needs to be
	// unmarshaled first.
	var payloadBytes []byte

	// Query the database for the sensor measurement. QueryRow because we expect
	// only one row to return. Scan maps to our go object(s) in the same order
	// as the query above.
	err := s.db.Pool.QueryRow(ctx, QUERY, deviceEUI).Scan(
		&measurement.DeviceEUI,
		&measurement.Timestamp,
		&payloadBytes,
		&measurement.CompanyID,
	)

	if err != nil {
		return domain.SensorMeasurement{}, fmt.Errorf("find latest: %w", err)
	}

	// Unmarshaling the JSONB bytes into the dynamic payload map.
	if err := json.Unmarshal(payloadBytes, &measurement.Payload); err != nil {
		return domain.SensorMeasurement{}, fmt.Errorf("unmarshal payload: %w", err)
	}

	return measurement, nil

}

// FindByTimeRange returns all measurements for a device within the given time window, ordered oldest first.
func (s *MeasurementRepository) FindByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]domain.SensorMeasurement, error) {

	// Query to fetch all the measurements from one device within a time range
	// ASC gives oldest first
	const QUERY = `
			SELECT device_eui, timestamp, payload, company_id
			FROM collection.sensor_measurement
			WHERE device_eui = $1
			  AND timestamp >= $2
			  AND timestamp <= $3
			ORDER BY timestamp ASC
	`
	// Query the database to collect all rows
	rows, err := s.db.Pool.Query(ctx, QUERY, deviceEUI, from, to)
	if err != nil {
		return nil, fmt.Errorf("find by time range: %w", err)
	}
	defer rows.Close()

	// The slice of sensormeasurements to be returned
	var measurements []domain.SensorMeasurement

	for rows.Next() {
		// The sensor measurement to be added to the slice
		var measurement domain.SensorMeasurement
		// JSONB object coming from postgres- Holds the raw JSONB bytes from postgres
		// pgx cannot scan JSONB directly into our map structure, so it needs to be
		// unmarshaled first.
		var payloadBytes []byte

		// Scan in the same column order as the query
		if err := rows.Scan(&measurement.DeviceEUI, &measurement.Timestamp, &payloadBytes, &measurement.CompanyID); err != nil {
			return nil, fmt.Errorf("scan measurement: %w", err)
		}

		// Unmarshaling the JSONB bytes into the dynamic payload map.
		if err := json.Unmarshal(payloadBytes, &measurement.Payload); err != nil {
			return nil, fmt.Errorf("unmarshal payload: %w", err)
		}

		// Add the measurement to the slice
		measurements = append(measurements, measurement)
	}

	// Check if the loop ended due to an error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return measurements, nil
}

// FindPayloadKeys retrieves unique payload keys from the last 10 readings for a device based on deviceEUI
func (r *MeasurementRepository) FindPayloadKeys (ctx context.Context, deviceEUI string) ([]payloadKeys string, error) {
	rows, err := r.db.Pool.Query(ctx, FindPayloadKeysQuery)
	if err != nil {
		return nil, fmt.Errorf("find payload keys: %w", err)
	}
	defer rows.Close()

	out := []string{}

	for rows.Next() {
		var key string
		err := rows.Scan(&key)
		if err != nil {
			return nil, fmt.Errorf("scan payload key: %w", err)
		}
		out = append(out, key)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return out, nil
}
