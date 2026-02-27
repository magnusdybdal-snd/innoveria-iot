package dto

type SensorResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	DeviceEUI  string `json:"device_eui"`
	GatewayEUI string `json:"gateway_eui"`
	Status     int    `json:"status"`
	LastSeenAt string `json:"lastSeenAt"`
}

type SensorListResponse struct {
	TotalCount int              `json:"total_count"`
	Sensors    []SensorResponse `json:"sensors"`
}
