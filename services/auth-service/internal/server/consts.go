package server

// Suppress godoclint here. Const naming should be explanatory in themselves.
//
//nolint:godoclint
const (
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	// main route
	AUTH_ROUTE = API_ROUTE + "/auth"

	// Company routes:
	COMPANY_ROUTE    = AUTH_ROUTE + "/companies"
	COMPANY_ID_ROUTE = COMPANY_ROUTE + "/{id}"

	// Factory routes:
	FACTORY_ROUTE    = AUTH_ROUTE + "/factories"
	FACTORY_ID_ROUTE = FACTORY_ROUTE + "/{id}"

	// Factory area routes:
	FACTORY_AREA_ROUTE    = AUTH_ROUTE + "/factory-areas"
	FACTORY_AREA_ID_ROUTE = FACTORY_AREA_ROUTE + "/{id}"

	// auth routes:
	LOGIN_ROUTE = AUTH_ROUTE + "/login"
	ME_ROUTE    = AUTH_ROUTE + "/me"
	// REGISTER_ROUTE = AUTH_ROUTE + "/register"
)
