package server

// Suppress godoclint in this file. Const naming should be explanatory in themselves
//godoclint:disable

const (
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	ERP_ROUTE    = API_ROUTE + "/erp"
	INGEST_ROUTE = ERP_ROUTE + "/ingest"

	ORDER_ROUTE            = INGEST_ROUTE + "/orders"
	ORDER_OPERATIONS_ROUTE = INGEST_ROUTE + "/order-operations"
	ORDER_REPORTINGS_ROUTE = INGEST_ROUTE + "/order-reportings"
	WORKCENTERS_ROUTE      = INGEST_ROUTE + "/workcenters"
)
