package server

// Suppress godoclint here. Const naming should be explanatory in themselves
//godoclint:disable

// API gateway routes
const (
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	// Proxy routing to correct service
	COLLECTION_ROUTE     = API_ROUTE + "/collection"
	DEVICE_ROUTE         = API_ROUTE + "/device"
	AUTHENTICATION_ROUTE = API_ROUTE + "/auth"
	CONTEXT_ROUTE        = API_ROUTE + "/context"
	ONBOARDING_ROUTE     = API_ROUTE + "/onboarding"
)
