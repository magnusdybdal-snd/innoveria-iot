package server

const (
	INDEX     = "/"
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	DEVICE_ROUTE = API_ROUTE + "/device"

	GATEWAY_ROUTE = DEVICE_ROUTE + "/gateways"
	SENSOR_ROUTE  = DEVICE_ROUTE + "/sensors"
)
