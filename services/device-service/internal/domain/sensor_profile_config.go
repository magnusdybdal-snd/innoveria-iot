package domain

import "context"

// SensorProfileConfig holds the configuration flags for a Chirpstack device profile.
// If no row exists for a profile, the default is ConfigurableSchema = false (fixed schema).
type SensorProfileConfig struct {
	ChirpstackProfileID string
	ConfigurableSchema  bool
}

// SensorProfileConfigRepository handles persistence of sensor profile configuration.
type SensorProfileConfigRepository interface {
	// Get returns the config for the given Chirpstack profile ID.
	// If no row exists, it returns a default config with ConfigurableSchema = false.
	Get(ctx context.Context, chirpstackProfileID string) (SensorProfileConfig, error)
	// Upsert inserts or updates the config for a Chirpstack profile.
	Upsert(ctx context.Context, sensorProfileCfg SensorProfileConfig) error
}

// SensorProfileConfigService defines the business logic for managing sensor profile configuration.
type SensorProfileConfigService interface {
	// Get returns the config for the given Chirpstack profile ID.
	// If no row exists, it returns a default config with ConfigurableSchema = false.
	Get(ctx context.Context, chirpstackProfileID string) (SensorProfileConfig, error)
	// Upsert inserts or updates the config for a Chirpstack profile.
	Upsert(ctx context.Context, sensorProfileCfg SensorProfileConfig) error
}
