package chirpstackrest

import "time"

// Chirpstack v4 model for Gateway
type ChirpstackGatewayList struct {
	Result     []ChirpstackGateway `json:"result"`
	TotalCount int                 `json:"totalCount"`
}

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

type ChirpstackGatewayLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Accuracy  float64 `json:"accuracy"`
	Source    string  `json:"source"`
}

// Chripstack post requests
type CreateChirpstackGatewayRequest struct {
	CreateGatewayPayload `json:"gateway"` // How chirpstack handles requests
}

type CreateGatewayPayload struct {
	GatewayEUI string `json:"gatewayId"` // chirpstack uses eui as id
	Name       string `json:"name"`
	TenantID   string `json:"tenantId"` // chirpstack uses tennant id, we use company id
}
