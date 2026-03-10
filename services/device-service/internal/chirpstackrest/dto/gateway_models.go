package dto

import "time"

// ChirpstackGatewayList is the paginated response from the Chirpstack GET /api/gateways endpoint.
type ChirpstackGatewayList struct {
	Result     []ChirpstackGateway `json:"result"`
	TotalCount int                 `json:"totalCount"`
}

// ChirpstackGateway represents a single gateway as returned by the Chirpstack API.
type ChirpstackGateway struct {
	GatewayEUI  string `json:"gatewayId"` // Chirpstack calls this gatewayId. We use our own id, so eui makes more sense
	Name        string `json:"name"`
	Description string `json:"description"`
	TenantID    string `json:"tenantId"`

	State string `json:"state"`

	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	LastSeenAt time.Time `json:"lastSeenAt"`

	Location   ChirpstackGatewayLocation `json:"location"`
	Properties map[string]string         `json:"properties"`
}

// ChirpstackGatewayLocation is the ChirpStack reported location of the gateway.
type ChirpstackGatewayLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Accuracy  float64 `json:"accuracy"`
	Source    string  `json:"source"`
}

// ChirpstackGatewayRequest is the top-level body for POST and PUT /api/gateways.
type ChirpstackGatewayRequest struct {
	GatewayPayload `json:"gateway"` // How chirpstack handles requests
}

// GatewayPayload contains the fields required by ChirpStack to register a new gateway.
type GatewayPayload struct {
	GatewayEUI string `json:"gatewayId"` // chirpstack uses eui as id
	Name       string `json:"name"`
	TenantID   string `json:"tenantId"` // chirpstack uses tennant id, we use company id
}
