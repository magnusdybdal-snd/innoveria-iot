package dto

// How we format the gateway response

/*
GET Request
*/
type GatewayResponse struct {
	ID         string `json:"id"`
	CompanyId  string `json:"company_id"`
	DeviceEUI  string `json:"device_eui"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
	LastSeenAt string `json:"lastSeenAt"`
}

type GatewayListResponse struct {
	TotalCount int               `json:"totalCount"`
	Gateways   []GatewayResponse `json:"gateways"`
}

/*
POST Request
*/
type CreateGatewayRequest struct {
	CompanyId  string `json:"company_id"`
	GatewayEUI string `json:"gateway_eui"`
	Name       string `json:"name"`
}
