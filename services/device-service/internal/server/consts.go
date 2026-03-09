// Package server provides HTTP server setup and route definitions for the device service.
package server

const (
	// INDEX is the root route path.
	INDEX = "/"
	// VERSION is the API version prefix.
	VERSION = "v1"
	// API_ROUTE is the base path for all API routes.
	API_ROUTE = "/api/" + VERSION

	// DEVICE_ROUTE is the base path for all device-related routes.
	DEVICE_ROUTE = API_ROUTE + "/device"

	// GATEWAY_ROUTE is the route for gateway collection endpoints.
	GATEWAY_ROUTE = DEVICE_ROUTE + "/gateways"
	// GATEWAY_ROUTE_ID is the route for single gateway endpoints, with an {id} path parameter.
	GATEWAY_ROUTE_ID = DEVICE_ROUTE + "/gateways/{id}"

	// SENSOR_ROUTE is the route for sensor collection endpoints.
	SENSOR_ROUTE = DEVICE_ROUTE + "/sensors"
	// SENSOR_ROUTE_ID is the route for single sensor endpoints, with an {id} path parameter.
	SENSOR_ROUTE_ID = DEVICE_ROUTE + "/sensors/{id}"

	// SENSOR_PROFILE_ROUTE is the route for sensor profile collection endpoints.
	SENSOR_PROFILE_ROUTE = DEVICE_ROUTE + "/sensor-profiles"
	// SENSOR_PROFILE_ROUTE_ID is the route for single sensor profile endpoints, with an {id} path parameter.
	SENSOR_PROFILE_ROUTE_ID = DEVICE_ROUTE + "/sensor-profiles/{id}"
)
