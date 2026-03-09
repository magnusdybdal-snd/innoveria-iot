package domain

// Status represents the runtime connectivity state of a device, as reported by Chirpstack.
type Status int

const (
	// StatusOnline indicates the device is currently connected and communicating.
	StatusOnline Status = iota
	// StatusNeverSeen indicates the device has never been observed by Chirpstack.
	StatusNeverSeen
	// StatusOffline indicates the device was previously seen but is no longer communicating.
	StatusOffline
)

// DeviceState represents the administrative state of a gateway or sensor stored
// in the database. Maps to the device.device_state enum in postgres.
type DeviceState string

const (
	// DeviceStateActive indicates the device is administratively enabled.
	DeviceStateActive DeviceState = "ACTIVE"
	// DeviceStateInactive indicates the device is administratively disabled.
	DeviceStateInactive DeviceState = "INACTIVE"
)
