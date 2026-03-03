// Package dto TODO(@vinjar): add proper documentation.
package dto

// GatewayResponse TODO(@vinjar): add proper documentation.
type GatewayResponse struct {
	ID         string `json:"id"`
	CompanyId  string `json:"company_id"`
	GatewayEUI string `json:"gateway_eui"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
	LastSeenAt string `json:"last_seen_at"`
}

// GatewayListResponse TODO(@vinjar): add proper documentation.
type GatewayListResponse struct {
	TotalCount int               `json:"total_count"`
	Gateways   []GatewayResponse `json:"gateways"`
}

// CreateGatewayRequest TODO(@vinjar): add proper documentation.
type CreateGatewayRequest struct {
	CompanyId  string `json:"company_id"`
	GatewayEUI string `json:"gateway_eui"`
	Name       string `json:"name"`
}
