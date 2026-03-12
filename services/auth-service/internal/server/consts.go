package server

// Suppress godoclint here. Const naming should be explanatory in themselves.
//
//godoclint:disable
const (
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	AUTH_ROUTE = API_ROUTE + "/auth"

	COMPANY_ROUTE    = AUTH_ROUTE + "/companies"
	COMPANY_ID_ROUTE = COMPANY_ROUTE + "/{id}"

	FACTORY_ROUTE    = AUTH_ROUTE + "/factories"
	FACTORY_ID_ROUTE = FACTORY_ROUTE + "/{id}"
)
