package dto

// SensorResponse TODO(@Magnus Dybdal): add proper documentation.
type SensorResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	DeviceEUI  string `json:"device_eui"`
	Status     int    `json:"status"`
	LastSeenAt string `json:"lastSeenAt"`
}

// SensorListResponse TODO(@Magnus Dybdal): add proper documentation.
type SensorListResponse struct {
	TotalCount int              `json:"total_count"`
	Sensors    []SensorResponse `json:"sensors"`
}

// UpdateSensorRequest represents the fields a caller can update on a sensor
type UpdateSensorRequest struct {
	Name                string `json:"name"`
	Description         string `json:"description"`
	FactoryAreaID       string `json:"factory_area_id"`
	ChirpstackProfileId string `json:"device_profile_id"`
}
