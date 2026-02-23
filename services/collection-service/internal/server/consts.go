package server

// Collection api routes
const (
	INDEX     = "/"
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	COLLECTION_ROUTE = API_ROUTE + "/collection" // For proxy routing

	// Routes:
	LATEST_MEASUREMENT = COLLECTION_ROUTE + "/latest"
)
