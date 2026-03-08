//godoclint:disable
package server

// Suppress godoclint in this file. Const naming should be explanatory in themselves
//godoclint:disable

const (
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	DEVICE_ROUTE = API_ROUTE + "/device"

	GATEWAY_ROUTE    = DEVICE_ROUTE + "/gateways"
	GATEWAY_ROUTE_ID = DEVICE_ROUTE + "/gateways/{id}"

	SENSOR_ROUTE    = DEVICE_ROUTE + "/sensors"
	SENSOR_ROUTE_ID = DEVICE_ROUTE + "/sensors/{id}"

	SENSOR_PROFILE_ROUTE    = DEVICE_ROUTE + "/sensor-profiles"
	SENSOR_PROFILE_ROUTE_ID = DEVICE_ROUTE + "/sensor-profiles/{id}"

	SENSOR_GROUP_ROUTE = DEVICE_ROUTE + "/sensor-groups"
)
