package server

const (
	INDEX     = "/"
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	DEVICE_ROUTE = API_ROUTE + "/device"

	GATEWAY_ROUTE    = DEVICE_ROUTE + "/gateways"
	GATEWAY_ROUTE_ID = DEVICE_ROUTE + "/gateways/{id}"
	SENSOR_ROUTE     = DEVICE_ROUTE + "/sensors"
)
