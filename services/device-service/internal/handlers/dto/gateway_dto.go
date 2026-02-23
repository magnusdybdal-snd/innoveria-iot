package dto

// How we format the gateway response
type GatewayResponse struct {
	ID         string `json:"id"`
	DeviceEUI  string `json:"device_eui"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
	LastSeenAt string `json:"lastSeenAt"`
}

type GatewayListResponse struct {
	TotalCount int               `json:"totalCount"`
	Gateways   []GatewayResponse `json:"gateways"`
}
