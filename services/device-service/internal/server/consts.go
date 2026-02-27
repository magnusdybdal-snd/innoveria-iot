package server

const (
	INDEX     = "/"
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	DEVICE_ROUTE = API_ROUTE + "/device"

	// Gateways routes
	GATEWAY_ROUTE    = DEVICE_ROUTE + "/gateways"
	GATEWAY_ROUTE_ID = DEVICE_ROUTE + "/gateways/{id}"

	// Sensors routes
	SENSOR_ROUTE = DEVICE_ROUTE + "/sensors"

	// Sensor profiles routes
	SENSOR_PROFILE_ROUTE    = DEVICE_ROUTE + "/sensor-profiles"
	SENSOR_PROFILE_ROUTE_ID = DEVICE_ROUTE + "/sensor-profiles/{id}"
)
