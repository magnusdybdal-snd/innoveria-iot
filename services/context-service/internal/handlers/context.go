package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers/dto"
	"innoveria-iot/pkg/authctx"
	"innoveria-iot/pkg/json"
)

// GetContextData computes context data for one or more sensors over a time window.
// @Summary		Get Context Data
// @Tags		context
// @Produce		json
// @Param		device_eui		query	[]string	true	"Device EUI(s)"	collectionFormat(multi)
// @Param		rule_id			query	string		true	"Aggregation Rule ID"
// @Param		from			query	string		true	"Period start (RFC3339)"
// @Param		to				query	string		true	"Period end (RFC3339)"
// @Param		bucket_minutes	query	int			false	"Bucket size in minutes (overrides rule default)"
// @Success		200	{array}		dto.ContextDataResponse
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Router		/data [get]
func GetContextData(svc domain.ContextService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		q := r.URL.Query()

		deviceEUIs := q["device_eui"]
		ruleID := q.Get("rule_id")
		fromStr := q.Get("from")
		toStr := q.Get("to")

		if len(deviceEUIs) == 0 {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing device_eui query parameter"), "at least one device_eui is required")
			return
		}
		if ruleID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing rule_id query parameter"), "rule_id is required")
			return
		}
		if fromStr == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing from query parameter"), "from is required")
			return
		}
		if toStr == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing to query parameter"), "to is required")
			return
		}

		from, err := time.Parse(time.RFC3339, fromStr)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid 'from' timestamp, expected RFC3339")
			return
		}
		to, err := time.Parse(time.RFC3339, toStr)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid 'to' timestamp, expected RFC3339")
			return
		}
		if !to.After(from) {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("'to' must be after 'from'"), "'to' must be after 'from'")
			return
		}

		bucketMins := 0
		if bStr := q.Get("bucket_minutes"); bStr != "" {
			bucketMins, err = strconv.Atoi(bStr)
			if err != nil || bucketMins <= 0 {
				json.HandleError(w, http.StatusBadRequest, fmt.Errorf("invalid bucket_minutes"), "bucket_minutes must be a positive integer")
				return
			}
		}

		results, err := svc.GetContextData(ctx, auth.CompanyID, deviceEUIs, ruleID, from, to, bucketMins)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "rule not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "failed to compute context data")
			return
		}

		if err := json.Encode(w, http.StatusOK, dto.MapContextDataDomainToDTO(results)); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}
