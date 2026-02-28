package dto

// Get requests
type SensorGroupResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
	Location  string `json:"location"`
}

type SensorGroupListResponse struct {
	TotalCount   int                   `json:"total_count"`
	SensorGroups []SensorGroupResponse `json:"sensor_groups"`
}

// Post requests
type CreateSensorGroup struct {
	CompanyId string `json:"company_id"`
	Name      string `json:"name"`
}
