package dto

// SensorResponse TODO(@Magnus Dybdal): add proper documentation.
type SensorResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	DeviceEUI  string `json:"device_eui"`
	Status     int    `json:"status"`
	LastSeenAt string `json:"last_seen_at"`
}

// SensorListResponse TODO(@Magnus Dybdal): add proper documentation.
type SensorListResponse struct {
	TotalCount int              `json:"total_count"`
	Sensors    []SensorResponse `json:"sensors"`
}

// UpdateSensorRequest represents the fields a caller can update on a sensor
type UpdateSensorRequest struct {
	Name                string  `json:"name"`
	Description         *string `json:"description"`
	FactoryAreaID       *string `json:"factory_area_id"`
	ChirpstackProfileID string  `json:"device_profile_id"`
	ProductionResource  *string `json:"production_resource"`
}

// CreateSensorRequest represents the fields required to register a new sensor.
type CreateSensorRequest struct {
	CompanyID           string  `json:"company_id"`
	Name                string  `json:"name"`
	Description         *string `json:"description"`
	DeviceEUI           string  `json:"device_eui"`
	ChirpstackProfileID string  `json:"device_profile_id"`
	FactoryAreaID       *string `json:"factory_area_id"`
	ProductionResource  *string `json:"production_resource"`
}
