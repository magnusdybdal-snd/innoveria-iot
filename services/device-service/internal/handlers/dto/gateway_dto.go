// Package dto contains all the data transfer objects and mappers used by the handlers.
package dto

// GatewayResponse represents a single gateway in API responses.
type GatewayResponse struct {
	ID            string  `json:"id"`
	CompanyID     string  `json:"company_id"`
	GatewayEUI    string  `json:"gateway_eui"`
	Name          string  `json:"name"`
	Description   *string `json:"description"`
	State         string  `json:"state"`
	FactoryAreaID *string `json:"factory_area_id"`
	Status        int     `json:"status"`
	LastSeenAt    string  `json:"last_seen_at"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// GatewayListResponse wraps a slice of GatewayResponse with a total count, returned by list endpoints.
type GatewayListResponse struct {
	TotalCount int               `json:"total_count"`
	Gateways   []GatewayResponse `json:"gateways"`
}

// CreateGatewayRequest contains the fields required to register a new gateway.
type CreateGatewayRequest struct {
	CompanyId  string `json:"company_id"    binding:"required"`
	GatewayEUI string `json:"gateway_eui"   binding:"required"`
	Name       string `json:"name"          binding:"required"`
}
