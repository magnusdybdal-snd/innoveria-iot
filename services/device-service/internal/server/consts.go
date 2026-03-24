// Package server provides HTTP server setup and route definitions for the device service.
package server

// Suppress godoclint in this file. Const naming should be explanatory in themselves
//godoclint:disable

const (
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

	// SENSOR_GROUP_ROUTE is the route for sensor group collection endpoints.
	SENSOR_GROUP_ROUTE = DEVICE_ROUTE + "/sensor-groups"

	// COMPANY_CONFIG_ROUTE is the route for company config endpoints.
	COMPANY_CONFIG_ROUTE = DEVICE_ROUTE + "/company-config"
	// COMPANY_CONFIG_ROUTE_ID is the route for single company config endpoints, with a {companyID} path parameter.
	COMPANY_CONFIG_ROUTE_ID = DEVICE_ROUTE + "/company-config/{companyID}"

	// MEASUREMENT_TYPE_ROUTE is the route for measurement type collection endpoints.
	MEASUREMENT_TYPE_ROUTE = DEVICE_ROUTE + "/measurement-types"
	// MEASUREMENT_TYPE_ROUTE_ALL is the route for listing all measurement types including deprecated.
	MEASUREMENT_TYPE_ROUTE_ALL = DEVICE_ROUTE + "/measurement-types/all"
	// MEASUREMENT_TYPE_ROUTE_DEPRECATE is the route for deprecating a single measurement type.
	MEASUREMENT_TYPE_ROUTE_DEPRECATE = DEVICE_ROUTE + "/measurement-types/{slug}/deprecate"
)
