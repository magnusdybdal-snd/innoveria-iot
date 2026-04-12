package handlers

import (
	"net/http"
)

// PostIngestOrderOperations handles batch ingest of manufacturing order operations
// pushed from erp-agent-service.
func PostIngestOrderOperations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r
		http.Error(w, "not implemented", http.StatusNotImplemented)
	}
}

// PostIngestOrderReportings handles batch ingest of manufacturing order operation
// reporting events pushed from erp-agent-service.
func PostIngestOrderReportings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r
		http.Error(w, "not implemented", http.StatusNotImplemented)
	}
}

// PostIngestWorkCenters handles batch ingest of work center master data pushed
// from erp-agent-service.
func PostIngestWorkCenters() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r
		http.Error(w, "not implemented", http.StatusNotImplemented)
	}
}
