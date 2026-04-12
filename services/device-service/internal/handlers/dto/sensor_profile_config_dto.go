package dto

// SensorProfileConfigResponse represents the configuration for a Chirpstack device profile in API responses.
type SensorProfileConfigResponse struct {
	ChirpstackProfileID string `json:"chirpstack_profile_id"`
	ConfigurableSchema  bool   `json:"configurable_schema"`
}

// PutSensorProfileConfigRequest represents the fields a caller can update on a sensor profile config.
type PutSensorProfileConfigRequest struct {
	ConfigurableSchema *bool `json:"configurable_schema"`
}
