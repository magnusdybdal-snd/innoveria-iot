package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"innoveria-iot/collection-service/internal/db"
	"innoveria-iot/collection-service/internal/domain"
	"time"
)

type SensorRepository struct {
	db *db.DB
}

func NewSensorRepository(db *db.DB) *SensorRepository {
	return &SensorRepository{db: db}
}

func (s *SensorRepository) Insert(ctx context.Context, measurement domain.SensorMeasurement, tenantID string) error {
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

func (s *SensorRepository) FindLatest(ctx context.Context, deviceEUI string) (domain.SensorMeasurement, error) {

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

func (s *SensorRepository) FindByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]domain.SensorMeasurement, error) {
	return nil, nil
}
