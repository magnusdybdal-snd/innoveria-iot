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
	getSensorProfileConfigQuery = `
		SELECT chirpstack_profile_id, configurable_schema
		FROM device.sensor_profile_config
		WHERE chirpstack_profile_id = $1
	`

	upsertSensorProfileConfigQuery = `
		INSERT INTO device.sensor_profile_config (chirpstack_profile_id, configurable_schema)
		VALUES ($1, $2)
		ON CONFLICT (chirpstack_profile_id) DO UPDATE
			SET configurable_schema = EXCLUDED.configurable_schema
	`
)

// SensorProfileConfigRepository
type SensorProfileConfigRepository struct {
	db *dbutil.DB
}

// NewSensorProfileConfigRepository
func NewSensorProfileConfigRepository(db *dbutil.DB) *SensorProfileConfigRepository {
	return &SensorProfileConfigRepository{db: db}
}

// Get implements SensorProfileConfigRepository.
func (r *SensorProfileConfigRepository) Get(ctx context.Context, chirpstackProfileID string) (domain.SensorProfileConfig, error) {
	var out domain.SensorProfileConfig
	err := r.db.Pool.QueryRow(ctx, getSensorProfileConfigQuery, chirpstackProfileID).Scan(
		&out.ChirpstackProfileID,
		&out.ConfigurableSchema,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.SensorProfileConfig{ChirpstackProfileID: chirpstackProfileID, ConfigurableSchema: false}, nil
		}
		return domain.SensorProfileConfig{}, fmt.Errorf("get sensor profile config: %w", err)
	}

	return out, nil
}

// Upsert implements SensorProfileConfigRepository.
func (r *SensorProfileConfigRepository) Upsert(ctx context.Context, sensorProfileCfg domain.SensorProfileConfig) error {
	_, err := r.db.Pool.Exec(ctx, upsertSensorProfileConfigQuery,
		sensorProfileCfg.ChirpstackProfileID,
		sensorProfileCfg.ConfigurableSchema,
	)
	if err != nil {
		return fmt.Errorf("upsert sensor profile config: %w", err)
	}

	return nil
}
