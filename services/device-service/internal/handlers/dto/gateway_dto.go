package dto

type GatewayResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
	LastSeenAt string `json:"lastSeenAt"`
}

type GatewayListResponse struct {
	TotalCount int               `json:"totalCount"`
	Gateways   []GatewayResponse `json:"gateways"`
}
