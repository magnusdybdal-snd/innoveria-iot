package server

// Suppress godoclint in this file. Const naming should be explanatory in themselves
//godoclint:disable

const (
	VERSION   = "v1"
	API_ROUTE = "/api/" + VERSION

	ERP_ROUTE    = API_ROUTE + "/erp"
	INGEST_ROUTE = ERP_ROUTE + "/ingest"

	// Ingest routes
	ORDER_ROUTE            = INGEST_ROUTE + "/orders"
	ORDER_OPERATIONS_ROUTE = INGEST_ROUTE + "/order-operations"
	ORDER_REPORTINGS_ROUTE = INGEST_ROUTE + "/order-reportings"
	WORKCENTERS_ROUTE      = INGEST_ROUTE + "/workcenters"

	// Extracting erp data
	PRODUCTION_RESOURCE_ROUTE    = ERP_ROUTE + "/production-resources"
	PRODUCTION_RESOURCE_ID_ROUTE = PRODUCTION_RESOURCE_ROUTE + "/{id}"

	ORDER_SUMMARY_ROUTE = ERP_ROUTE + "/order"
)
