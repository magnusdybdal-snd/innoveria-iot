package dto

import "time"

// ChirpstackSensorList TODO(@vinjar): add proper documentation.
type ChirpstackSensorList struct {
	Result     []ChirpstackSensor `json:"result"`
	TotalCount int                `json:"totalCount"`
}

// ChirpstackSensor TODO(@vinjar): add proper documentation.
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

// ChirpstackSensorStatus TODO(@vinjar): add proper documentation.
type ChirpstackSensorStatus struct {
	BatteryLevel        float64 `json:"batteryLevel"`
	ExternalPowerSource bool    `json:"externalPowerSource"`
	Margin              int     `json:"margin"`
}
