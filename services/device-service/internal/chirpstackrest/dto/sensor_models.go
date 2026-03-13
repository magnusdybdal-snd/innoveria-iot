package dto

import "time"

// ChirpstackSensorList is the paginated response from the ChirpStack GET /api/devices endpoint.
type ChirpstackSensorList struct {
	Result     []ChirpstackSensor `json:"result"`
	TotalCount int                `json:"totalCount"`
}

// ChirpstackSensor represents a single device as returned by the ChirpStack API.
// ChirpStack refers to sensors as "devices" internally.
type ChirpstackSensor struct {
	DeviceEUI         string `json:"devEui"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DeviceProfileID   string `json:"deviceProfileId"`
	DeviceProfileName string `json:"deviceProfileName"`

	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	LastSeenAt time.Time `json:"lastSeenAt"`

	DeviceStatus ChirpstackSensorStatus `json:"deviceStatus"`
	Tags         map[string]string      `json:"tags"`
}

// ChirpstackSensorStatus contains hardware status reported by the device to ChirpStack.
type ChirpstackSensorStatus struct {
	BatteryLevel        float64 `json:"batteryLevel"`
	ExternalPowerSource bool    `json:"externalPowerSource"`
	Margin              int     `json:"margin"`
}

// ChirpstackSensorRequest is the top-level body for POST and PUT /api/devices.
// ChirpStack expects the payload wrapped under a "device" key.
type ChirpstackSensorRequest struct {
	SensorPayload `json:"device"`
}

// SensorPayload contains the fields required by ChirpStack to register a new device.
type SensorPayload struct {
	DeviceEUI       string `json:"devEui"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ApplicationID   string `json:"applicationId"`
	DeviceProfileID string `json:"deviceProfileId"`
	JoinEUI         string `json:"joinEui"`
}

// ChirpstackSensorKeyRequest is the body for POST /api/devices/{devEui}/keys.
// Chirpstack expects the request wrapped under a "deviceKeys" key
type ChirpstackSensorKeyRequest struct {
	DeviceKeys SensorKeysPayload `json:"deviceKeys"`
}

// SensorKeysPayload contains the OTAA root key for a device.
type SensorKeysPayload struct {
	DevEUI string `json:"devEui"`
	NwkKey string `json:"nwkKey"`
}
