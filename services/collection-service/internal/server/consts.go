// Package server provides the HTTP server and routing for the device service API.
package server

// Suppress godoclint in this file. Const naming should be explanatory in themselves
//godoclint:disable

// Collection api routes
const (
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	COLLECTION_ROUTE = API_ROUTE + "/collection" // For proxy routing

	// Routes:
	LATEST_MEASUREMENT   = COLLECTION_ROUTE + "/latest"
	MEASUREMENTS_BY_TIME = COLLECTION_ROUTE + "/measurements"
	PAYLOAD_TAGS         = COLLECTION_ROUTE + "/payload-tags"

	COMPANY_CONFIG_ROUTE = COLLECTION_ROUTE + "/company-config"
)
