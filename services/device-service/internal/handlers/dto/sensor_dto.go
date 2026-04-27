package dto

// SampleEUIResponse represents the response for the sample EUI endpoint.
type SampleEUIResponse struct {
	DeviceEUI string `json:"device_eui"`
}

// SensorResponse represents a single sensor in API responses.
type SensorResponse struct {
	ID                  string  `json:"id"`
	CompanyID           string  `json:"company_id"`
	Name                string  `json:"name"`
	Description         *string `json:"description"`
	ElectricitySensor   bool    `json:"electricity_sensor"`
	Voltage             *int    `json:"voltage"`
	DeviceEUI           string  `json:"device_eui"`
	AppKey              string  `json:"app_key"`
	State               string  `json:"state"`
	FactoryID           string  `json:"factory_id"`
	FactoryAreaID       string  `json:"factory_area_id"`
	ProductionResource  *int64  `json:"production_resource"`
	ChirpstackProfileID string  `json:"device_profile_id"`
	Status              int     `json:"status"`
	LastSeenAt          string  `json:"last_seen_at"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
}

// SensorListResponse wraps a slice of SensorResponse with a total count, returned by list endpoints.
type SensorListResponse struct {
	TotalCount int              `json:"total_count"`
	Sensors    []SensorResponse `json:"sensors"`
}

// UpdateSensorRequest represents the fields a caller can update on a sensor
type UpdateSensorRequest struct {
	Name                *string `json:"name"`
	ChirpstackProfileID *string `json:"device_profile_id"`
	Description         *string `json:"description"`
	ElectricitySensor   *bool   `json:"electricity_sensor"`
	Voltage             *int    `json:"voltage"`
	FactoryID           *string `json:"factory_id"`
	FactoryAreaID       *string `json:"factory_area_id"`     // Optional — UUID, omit if service not yet available
	ProductionResource  *int64  `json:"production_resource"` // Optional — int64 ERP production resource ID
}

// CreateSensorRequest represents the fields required to register a new sensor.
type CreateSensorRequest struct {
	Name                string  `json:"name"                  binding:"required"`
	DeviceEUI           string  `json:"device_eui"            binding:"required"`
	AppKey              string  `json:"app_key"               binding:"required"`
	ChirpstackProfileID string  `json:"device_profile_id"     binding:"required"`
	Description         *string `json:"description"`
	ElectricitySensor   bool    `json:"electricity_sensor"`
	Voltage             *int    `json:"voltage"`
	FactoryID           string  `json:"factory_id"            binding:"required"`
	FactoryAreaID       string  `json:"factory_area_id"       binding:"required"`
	ProductionResource  *int64  `json:"production_resource"` // Optional — int64 ERP production resource ID
}
