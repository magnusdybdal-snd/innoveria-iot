package dto

// SensorGroupResponse represents a single sensor group in API responses.
type SensorGroupResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
	Location  string `json:"location"`
}

// SensorGroupListResponse wraps a slice of SensorGroupResponse with a total count, returned by list endpoints.
type SensorGroupListResponse struct {
	TotalCount   int                   `json:"total_count"`
	SensorGroups []SensorGroupResponse `json:"sensor_groups"`
}

// CreateSensorGroup contains the fields required to create a new sensor group.
type CreateSensorGroup struct {
	CompanyId string `json:"company_id"`
	Name      string `json:"name"`
}
