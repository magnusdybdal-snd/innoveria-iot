// Package server TODO(@vinjar): add proper documentation.
package server

const (
	// INDEX TODO(@vinjar): add proper documentation.
	INDEX = "/"
	// VERSION TODO(@vinjar): add proper documentation.
	VERSION = "v1"
	// API_ROUTE TODO(@vinjar): add proper documentation.
	API_ROUTE = "/api/" + VERSION

	// DEVICE_ROUTE TODO(@vinjar): add proper documentation.
	DEVICE_ROUTE = API_ROUTE + "/device"

	// GATEWAY_ROUTE TODO(@vinjar): add proper documentation.
	GATEWAY_ROUTE = DEVICE_ROUTE + "/gateways"
	// GATEWAY_ROUTE_ID TODO(@vinjar): add proper documentation.
	GATEWAY_ROUTE_ID = DEVICE_ROUTE + "/gateways/{id}"

	// SENSOR_ROUTE TODO(@vinjar): add proper documentation.
	SENSOR_ROUTE = DEVICE_ROUTE + "/sensors"
	// SENSOR_ROUTE_ID TODO(@vinjar): add proper documentation.
	SENSOR_ROUTE_ID = DEVICE_ROUTE + "/sensors/{id}"

	// SENSOR_PROFILE_ROUTE TODO(@vinjar): add proper documentation.
	SENSOR_PROFILE_ROUTE = DEVICE_ROUTE + "/sensor-profiles"
	// SENSOR_PROFILE_ROUTE_ID TODO(@vinjar): add proper documentation.
	SENSOR_PROFILE_ROUTE_ID = DEVICE_ROUTE + "/sensor-profiles/{id}"

	// SENSOR_GROUP_ROUTE TODO(@vinjar): add proper documentation.
	SENSOR_GROUP_ROUTE = DEVICE_ROUTE + "/sensor-groups"
)
