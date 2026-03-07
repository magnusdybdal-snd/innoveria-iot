package server

// Suppress godoclint here. Const naming should be explanatory in themselves.
//
//godoclint:disable
const (
	INDEX     = "/"
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	AUTH_ROUTE    = API_ROUTE + "/auth"
	COMPANY_ROUTE = AUTH_ROUTE + "/company"
)
