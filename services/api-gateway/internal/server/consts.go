package server

// API gateway routes 
const (
	INDEX = "/"
	VERSION = "v1"
	API_ROUTE = "/api/" + VERSION 
	
	COLLECTION_ROUTE = API_ROUTE + "/collection"
	AUTHENTICATION_ROUTE = API_ROUTE + "/authentication"
	CONTEXT_ROUTE = API_ROUTE + "/context"
)
