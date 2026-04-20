package repository

import (
	"context"
	"errors"
	"fmt"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5"
)

const (
	createSensorQuery = `
		INSERT INTO device.sensor (company_id, device_eui, app_key, name, description, state, factory_id, factory_area_id, production_resource_id, chirpstack_profile_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING sensor_id, company_id, device_eui, app_key, name, description, state, factory_id, factory_area_id, production_resource_id, chirpstack_profile_id, created_at, updated_at
	`

	findSensorByIDQuery = `
		SELECT sensor_id, company_id, device_eui, app_key, name, description, state, factory_id, factory_area_id, production_resource_id, chirpstack_profile_id, created_at, updated_at
		FROM device.sensor
		WHERE sensor_id = $1
	`

	findAllSensorsByCompanyIDQuery = `
		SELECT sensor_id, company_id, device_eui, app_key, name, description, state, factory_id, factory_area_id, production_resource_id, chirpstack_profile_id, created_at, updated_at
		FROM device.sensor
		WHERE company_id = $1
		ORDER BY created_at ASC
	`

	findByProductionResourceIDQuery = `
		SELECT sensor_id, company_id, device_eui, app_key, name, description, state, factory_id, factory_area_id, production_resource_id, chirpstack_profile_id, created_at, updated_at
		FROM device.sensor
		WHERE production_resource_id = $1 AND company_id = $2
		ORDER BY created_at ASC
	`

	findSensorByEUIQuery = `
		SELECT sensor_id, company_id, device_eui, app_key, name, description, state, factory_id, factory_area_id, production_resource_id, chirpstack_profile_id, created_at, updated_at
		FROM device.sensor
		WHERE device_eui = $1
	`

	updateSensorStateQuery = `
		UPDATE device.sensor
		SET state = $1, updated_at = now()
		WHERE sensor_id = $2
	`

	updateSensorQuery = `
		UPDATE device.sensor
		SET name = $1, description = $2, factory_id = $3, factory_area_id = $4, chirpstack_profile_id = $5, updated_at = now()
		WHERE sensor_id = $6
	`

	deleteSensorQuery = `
		DELETE FROM device.sensor
		WHERE sensor_id = $1
	`
)

// SensorRepository handles persistence of sensor metadata stored in the database.
type SensorRepository struct {
	db *dbutil.DB
}

// NewSensorRepository creates a new SensorRepository with the given database.
func NewSensorRepository(db *dbutil.DB) *SensorRepository {
	return &SensorRepository{db: db}
}

// Create inserts a new sensor into the database.
// Returns the newly created sensor with database generated fields.
func (r *SensorRepository) Create(ctx context.Context, sensor domain.Sensor) (domain.Sensor, error) {

	var out domain.Sensor
	err := r.db.Pool.QueryRow(ctx, createSensorQuery,
		sensor.CompanyID,
		sensor.DeviceEUI,
		sensor.AppKey,
		sensor.Name,
		sensor.Description,
		sensor.State,
		sensor.FactoryID,
		sensor.FactoryAreaID,
		sensor.ProductionResource,
		sensor.ChirpstackProfileID,
	).Scan(
		&out.Id,
		&out.CompanyID,
		&out.DeviceEUI,
		&out.AppKey,
		&out.Name,
		&out.Description,
		&out.State,
		&out.FactoryID,
		&out.FactoryAreaID,
		&out.ProductionResource,
		&out.ChirpstackProfileID,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		return domain.Sensor{}, fmt.Errorf("create sensor: %w", err)
	}

	return out, nil
}

// FindByID retrieves a sensor by its internal UUID.
func (r *SensorRepository) FindByID(ctx context.Context, sensorID string) (domain.Sensor, error) {

	var out domain.Sensor
	err := r.db.Pool.QueryRow(ctx, findSensorByIDQuery, sensorID).Scan(
		&out.Id,
		&out.CompanyID,
		&out.DeviceEUI,
		&out.AppKey,
		&out.Name,
		&out.Description,
		&out.State,
		&out.FactoryID,
		&out.FactoryAreaID,
		&out.ProductionResource,
		&out.ChirpstackProfileID,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		return domain.Sensor{}, fmt.Errorf("find sensor by sensor id: %w", err)
	}

	return out, nil

}

// FindAllByCompanyID retrieves all sensors belonging to the given company, ordered by creation date.
func (r *SensorRepository) FindAllByCompanyID(ctx context.Context, companyID string) ([]domain.Sensor, error) {

	// Query the database to collect all rows
	rows, err := r.db.Pool.Query(ctx, findAllSensorsByCompanyIDQuery, companyID)
	if err != nil {
		return nil, fmt.Errorf("find all sensors by company id: %w", err)
	}
	defer rows.Close()

	// The slice of sensors to be returned
	var out []domain.Sensor

	for rows.Next() {
		var sensor domain.Sensor

		err := rows.Scan(
			&sensor.Id,
			&sensor.CompanyID,
			&sensor.DeviceEUI,
			&sensor.AppKey,
			&sensor.Name,
			&sensor.Description,
			&sensor.State,
			&sensor.FactoryID,
			&sensor.FactoryAreaID,
			&sensor.ProductionResource,
			&sensor.ChirpstackProfileID,
			&sensor.CreatedAt,
			&sensor.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan sensor: %w", err)
		}

		// Add the sensor to the slice
		out = append(out, sensor)
	}

	// Sanity check if the loop ended due to an error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return out, nil

}

// FindByProductionResourceID retrieves sensor by their production resource id. returns an empty slice if no sensors found on that resource.
func (r *SensorRepository) FindByProductionResourceID(ctx context.Context, companyID string, productionResourceID string) ([]domain.Sensor, error) {

	rows, err := r.db.Pool.Query(ctx, findByProductionResourceIDQuery, productionResourceID, companyID)
	if err != nil {
		return nil, fmt.Errorf("find sensors by production resource id: %w", err)
	}
	defer rows.Close()

	out := []domain.Sensor{}

	for rows.Next() {
		var sensor domain.Sensor

		err := rows.Scan(
			&sensor.Id,
			&sensor.CompanyID,
			&sensor.DeviceEUI,
			&sensor.AppKey,
			&sensor.Name,
			&sensor.Description,
			&sensor.State,
			&sensor.FactoryID,
			&sensor.FactoryAreaID,
			&sensor.ProductionResource,
			&sensor.ChirpstackProfileID,
			&sensor.CreatedAt,
			&sensor.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan sensor: %w", err)
		}

		out = append(out, sensor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return out, nil

}

// FindByEUI retrieves a sensor by its hardware EUI. Used for deduplication checks before registration.
func (r *SensorRepository) FindByEUI(ctx context.Context, deviceEUI string) (domain.Sensor, error) {

	var out domain.Sensor
	err := r.db.Pool.QueryRow(ctx, findSensorByEUIQuery, deviceEUI).Scan(
		&out.Id,
		&out.CompanyID,
		&out.DeviceEUI,
		&out.AppKey,
		&out.Name,
		&out.Description,
		&out.State,
		&out.FactoryID,
		&out.FactoryAreaID,
		&out.ProductionResource,
		&out.ChirpstackProfileID,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Sensor{}, domain.ErrNotFound
		}
		return domain.Sensor{}, fmt.Errorf("find sensor by eui: %w", err)
	}

	return out, nil

}

// UpdateState sets the administrative state of a sensor and updates the updated at timestamp.
// Returns an error if no sensor with the given ID exists.
func (r *SensorRepository) UpdateState(ctx context.Context, sensorID string, state domain.DeviceState) error {

	tag, err := r.db.Pool.Exec(ctx, updateSensorStateQuery, state, sensorID)
	if err != nil {
		return fmt.Errorf("update sensor state: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("sensor not found: %s", sensorID)
	}

	return nil
}

// Update updates the editable fields of a sensor (name, description, factory area, Chirpstack profile) and refreshes the updated at timestamp.
// Returns an error if no sensor with the given ID exists.
func (r *SensorRepository) Update(ctx context.Context, sensorID string, payload domain.Sensor) error {

	tag, err := r.db.Pool.Exec(ctx, updateSensorQuery,
		payload.Name,
		payload.Description,
		payload.FactoryID,
		payload.FactoryAreaID,
		payload.ChirpstackProfileID,
		sensorID,
	)

	if err != nil {
		return fmt.Errorf("update sensor %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("sensor not found: %s", sensorID)
	}

	return nil
}

// Delete tries to delete a sensor from the database.
// Returns an error if deletion fails or no sensor is found.
func (r *SensorRepository) Delete(ctx context.Context, deviceID string) error {

	tag, err := r.db.Pool.Exec(ctx, deleteSensorQuery, deviceID)
	if err != nil {
		return fmt.Errorf("delete sensor %s: %w", deviceID, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("sensor not found: %s", deviceID)
	}

	return nil
}
