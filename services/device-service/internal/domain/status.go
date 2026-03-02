package domain

type Status int

const (
	StatusOnline Status = iota
	StatusNeverSeen
	StatusOffline
)

// DeviceState represents the administrative state of a gateway or sensor stored
// in the database. Maps to the device.device_state enum in postgres.
type DeviceState string

const (
	DeviceStateActive   DeviceState = "ACTIVE"
	DeviceStateInactive DeviceState = "INACTIVE"
)
