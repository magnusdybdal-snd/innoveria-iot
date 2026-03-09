package server

const (
	// VERSION is the API version prefix.
	VERSION = "v1"
	// API_ROUTE is the base path for all API routes.
	API_ROUTE = "/api/" + VERSION

	// COMPANY_ROUTE is the route for company onboarding endpoints
	COMPANY_ROUTE = API_ROUTE + "/company"
)
