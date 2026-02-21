package chirpstackrest

import "time"

// Chirpstack v4 model for Device (Sensor)
type ChirpstackSensorList struct {
	Result     []ChirpstackSensor `json:"result"`
	TotalCount int                `json:"totalCount"`
}

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

type ChirpstackSensorStatus struct {
	BatteryLevel        float64 `json:"batteryLevel"`
	ExternalPowerSource bool    `json:"externalPowerSource"`
	Margin              int     `json:"margin"`
}
