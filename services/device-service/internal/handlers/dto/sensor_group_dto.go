package dto

// SensorGroupResponse TODO(@vinjar): add proper documentation.
type SensorGroupResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
	Location  string `json:"location"`
}

// SensorGroupListResponse TODO(@vinjar): add proper documentation.
type SensorGroupListResponse struct {
	TotalCount   int                   `json:"total_count"`
	SensorGroups []SensorGroupResponse `json:"sensor_groups"`
}

// CreateSensorGroup TODO(@vinjar): add proper documentation.
type CreateSensorGroup struct {
	CompanyId string `json:"company_id"`
	Name      string `json:"name"`
}
