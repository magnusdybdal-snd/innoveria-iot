package domain

// Status TODO(@Magnus Dybdal): add proper documentation.
type Status int

const (
	// StatusOnline TODO(@Magnus Dybdal): add proper documentation.
	StatusOnline Status = iota
	// StatusNeverSeen TODO(@Magnus Dybdal): add proper documentation.
	StatusNeverSeen
	// StatusOffline TODO(@Magnus Dybdal): add proper documentation.
	StatusOffline
)

// DeviceState represents the administrative state of a gateway or sensor stored
// in the database. Maps to the device.device_state enum in postgres.
type DeviceState string

const (
	// DeviceStateActive TODO(@Magnus Dybdal): add proper documentation.
	DeviceStateActive DeviceState = "ACTIVE"
	// DeviceStateInactive TODO(@Magnus Dybdal): add proper documentation.
	DeviceStateInactive DeviceState = "INACTIVE"
)
