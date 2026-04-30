package handlers

import (
	"errors"
	"net/http"
	"strings"
)

// companyIDHeader is the trusted company id header injected by middleware.
const companyIDHeader = "X-Erp-Company-Id"

// companyIDFromHeader reads and validates company id from request headers.
// It returns an error if X-Erp-Company-Id is missing or blank.
func companyIDFromHeader(r *http.Request) (string, error) {
	companyID := strings.TrimSpace(r.Header.Get(companyIDHeader))
	if companyID == "" {
		return "", errors.New("missing company id header")
	}
	return companyID, nil
}
