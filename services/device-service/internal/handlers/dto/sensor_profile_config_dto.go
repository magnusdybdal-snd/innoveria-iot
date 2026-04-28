package dto

// SensorProfileConfigResponse represents the configuration for a Chirpstack device profile in API responses.
type SensorProfileConfigResponse struct {
	ChirpstackProfileID string `json:"chirpstack_profile_id"`
	ConfigurableSchema  bool   `json:"configurable_schema"`
}

// PutSensorProfileConfigRequest represents the fields a caller can set on a sensor profile config.
type PutSensorProfileConfigRequest struct {
	// ConfigurableSchema uses *bool to distinguish an explicit false from a missing field,
	// since Go's JSON decoder cannot differentiate the two for plain bool types.
	// The field is required — a nil value is rejected with 400.
	ConfigurableSchema *bool `json:"configurable_schema"`
}
