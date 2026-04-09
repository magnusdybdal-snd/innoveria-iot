// Package server provides the HTTP server and routing for the context service API.
package server

// Suppress godoclint in this file. Const naming should be explanatory in themselves
//godoclint:disable

// Context service API routes
const (
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	CONTEXT_ROUTE = API_ROUTE + "/context"

	// Routes:
	RULES_ROUTE        = CONTEXT_ROUTE + "/rules"
	RULES_ID_ROUTE     = RULES_ROUTE + "/{id}"
	CONTEXT_DATA_ROUTE = CONTEXT_ROUTE + "/data"
)
