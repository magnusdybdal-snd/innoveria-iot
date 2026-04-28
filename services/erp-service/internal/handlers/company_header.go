package handlers

import (
	"errors"
	"net/http"
	"strings"
)

const (
	// companyIDHeader is injected by RequireAgentAuth for ingest routes.
	companyIDHeader = "X-Erp-Company-Id"
	// authCompanyIDHeader is injected by the api-gateway for user-facing routes.
	authCompanyIDHeader = "X-Auth-Company-Id"
)

// companyIDFromHeader reads X-Erp-Company-Id, injected by RequireAgentAuth on ingest routes.
func companyIDFromHeader(r *http.Request) (string, error) {
	companyID := strings.TrimSpace(r.Header.Get(companyIDHeader))
	if companyID == "" {
		return "", errors.New("missing company id header")
	}
	return companyID, nil
}

// authCompanyIDFromHeader reads X-Auth-Company-Id, injected by the api-gateway on user-facing routes.
func authCompanyIDFromHeader(r *http.Request) (string, error) {
	companyID := strings.TrimSpace(r.Header.Get(authCompanyIDHeader))
	if companyID == "" {
		return "", errors.New("missing company id header")
	}
	return companyID, nil
}
