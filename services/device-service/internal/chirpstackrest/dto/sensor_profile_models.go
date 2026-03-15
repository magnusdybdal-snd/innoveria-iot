package dto

import "time"

// DeviceProfileListResponse is the paginated response from the Chirpstack GET /api/device-profiles endpoint.
type DeviceProfileListResponse struct {
	Result     []DeviceProfile `json:"result"`
	TotalCount int             `json:"totalCount"`
}

// DeviceProfile represents a single LoRaWAN device profile as returned by the Chirpstack API.
type DeviceProfile struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	DeviceID          string    `json:"deviceId"`
	DeviceName        string    `json:"deviceName"`
	FirmwareVersion   string    `json:"firmwareVersion"`
	MACVersion        string    `json:"macVersion"`
	RegParamsRevision string    `json:"regParamsRevision"`
	Region            string    `json:"region"`
	SupportsClassB    bool      `json:"supportsClassB"`
	SupportsClassC    bool      `json:"supportsClassC"`
	SupportsOtaa      bool      `json:"supportsOtaa"`
	VendorID          string    `json:"vendorId"`
	VendorName        string    `json:"vendorName"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

// DeviceProfileDetail is the full device profile returned by GET /api/device-profiles/{id}.
// Global profiles (created by import-device-profiles) have an empty TenantID.
type DeviceProfileDetail struct {
	ID                      string            `json:"id"`
	TenantID                string            `json:"tenantId"`
	Name                    string            `json:"name"`
	Description             string            `json:"description"`
	Region                  string            `json:"region"`
	MACVersion              string            `json:"macVersion"`
	RegParamsRevision       string            `json:"regParamsRevision"`
	ADRAlgorithmID          string            `json:"adrAlgorithmId"`
	PayloadCodecRuntime     string            `json:"payloadCodecRuntime"`
	PayloadCodecScript      string            `json:"payloadCodecScript"`
	FlushQueueOnActivate    bool              `json:"flushQueueOnActivate"`
	UplinkInterval          int               `json:"uplinkInterval"`
	DeviceStatusReqInterval int               `json:"deviceStatusReqInterval"`
	SupportsOtaa            bool              `json:"supportsOtaa"`
	SupportsClassB          bool              `json:"supportsClassB"`
	SupportsClassC          bool              `json:"supportsClassC"`
	Tags                    map[string]string `json:"tags"`
}

// DeviceProfileDetailResponse wraps the full profile in the Chirpstack response envelope.
type DeviceProfileDetailResponse struct {
	DeviceProfile DeviceProfileDetail `json:"deviceProfile"`
}

// CreateDeviceProfileRequest is the request body for POST /api/device-profiles.
type CreateDeviceProfileRequest struct {
	DeviceProfile CreateDeviceProfileBody `json:"deviceProfile"`
}

// CreateDeviceProfileBody holds the fields for creating a tenant-level device profile.
type CreateDeviceProfileBody struct {
	TenantID                string            `json:"tenantId"`
	Name                    string            `json:"name"`
	Description             string            `json:"description"`
	Region                  string            `json:"region"`
	MACVersion              string            `json:"macVersion"`
	RegParamsRevision       string            `json:"regParamsRevision"`
	ADRAlgorithmID          string            `json:"adrAlgorithmId"`
	PayloadCodecRuntime     string            `json:"payloadCodecRuntime"`
	PayloadCodecScript      string            `json:"payloadCodecScript"`
	FlushQueueOnActivate    bool              `json:"flushQueueOnActivate"`
	UplinkInterval          int               `json:"uplinkInterval"`
	DeviceStatusReqInterval int               `json:"deviceStatusReqInterval"`
	SupportsOtaa            bool              `json:"supportsOtaa"`
	SupportsClassB          bool              `json:"supportsClassB"`
	SupportsClassC          bool              `json:"supportsClassC"`
	Tags                    map[string]string `json:"tags"`
}

// CreateDeviceProfileResponse is the response from POST /api/device-profiles.
type CreateDeviceProfileResponse struct {
	ID string `json:"id"`
}
