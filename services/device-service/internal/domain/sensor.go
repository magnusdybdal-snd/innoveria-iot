package domain

import "time"

type Sensor struct {
	DevEUI     string
	GatewayEUI string
	LastSeenAt time.Time
}
