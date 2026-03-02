package server

// API gateway routes
const (
	INDEX     = "/"
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	// Proxy routing to correct service
	COLLECTION_ROUTE     = API_ROUTE + "/collection"
	DEVICE_ROUTE         = API_ROUTE + "/device"
	AUTHENTICATION_ROUTE = API_ROUTE + "/authentication"
	CONTEXT_ROUTE        = API_ROUTE + "/context"
)
