package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"

	"github.com/google/uuid"
)

// GetRules retrieves all aggregation rules for a given company ID.
// @Summary 		Get Aggregation Rules
// @Tags 			context
// @Produce 		json
// @Param 			company_id query string true "Company ID"
// @Success 		200 {array} dto.AggregationRuleResponse
// @Failure 		400
// @Failure 		500
// @Router 			/rules [get]
func GetRules(svc domain.RuleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		companyID := r.URL.Query().Get("company_id")

		if companyID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing company_id query parameter"), "company_id is required")
			return
		}

		rules, err := svc.GetRules(ctx, companyID)
		if err != nil {
			if errors.Is(err, domain.ErrDatabase) {
				json.HandleError(w, http.StatusInternalServerError, err, "database error")
			} else {
				json.HandleError(w, http.StatusInternalServerError, err, "failed to retrieve rules")
			}
			return
		}

		response := make([]dto.AggregationRuleResponse, len(rules))
		for i, rule := range rules {
			response[i] = dto.AggregationRuleResponse{
				ID:                rule.ID,
				CompanyID:         rule.CompanyID,
				Name:              rule.Name,
				ContextType:       rule.ContextType,
				MeasurementType:   rule.MeasurementType,
				AggregationMethod: string(rule.AggregationMethod),
				TimeBucketMinutes: rule.TimeBucketMinutes,
				IsActive:          rule.IsActive,
				CreatedAt:         rule.CreatedAt,
				UpdatedAt:         rule.UpdatedAt,
			}
		}

		if err := json.Encode(w, http.StatusOK, response); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}

// CreateRule creates a new aggregation rule in the database and returns its ID.
// @Summary 		Create Aggregation Rule
// @Tags 			context
// @Accept 			json
// @Produce 		json
// @Param 			rule body dto.CreateAggregationRuleRequest true "Aggregation Rule to create"
// @Success 		201
// @Failure 		400
// @Failure 		409
// @Failure 		500
// @Router 			/rules [post]
func CreateRule(svc domain.RuleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req, err := json.Decode[dto.CreateAggregationRuleRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid request body")
			return
		}

		if _, err := uuid.Parse(req.CompanyID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid company_id (uuid)")
			return
		}

		if req.Name == "" || req.ContextType == "" || req.MeasurementType == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing required fields"), "name, context_type and measurement_type are required")
			return
		}

		if !domain.AggMethod(req.AggregationMethod).IsValid() {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("invalid aggregation_method: %s", req.AggregationMethod), "aggregation_method must be one of: AVG, SUM, MAX, MIN")
			return
		}

		if req.TimeBucketMinutes <= 0 {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("invalid time_bucket_minutes: %d", req.TimeBucketMinutes), "time interval must be greater than 0")
			return
		}

		rule := domain.AggregationRule{
			CompanyID:         req.CompanyID,
			Name:              req.Name,
			ContextType:       req.ContextType,
			MeasurementType:   req.MeasurementType,
			AggregationMethod: domain.AggMethod(req.AggregationMethod),
			TimeBucketMinutes: req.TimeBucketMinutes,
			IsActive:          req.IsActive,
		}

		ruleID, err := svc.CreateRule(ctx, rule)
		if err != nil {
			if errors.Is(err, domain.ErrConflict) {
				json.HandleError(w, http.StatusConflict, err, "a rule with this context_type already exists for the company")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "failed to create rule")
			return
		}

		if err := json.Encode(w, http.StatusCreated, ruleID); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}

// DeleteRule deletes an aggregation rule by its ID.
//
// @Summary 		Delete Aggregation Rule
// @Tags 			context
// @Param 			id path string true "id"
// @Success 		204
// @Failure 		400
// @Failure 		404
// @Failure 		500
// @Router 			/rules/{id} [delete]
func DeleteRule(svc domain.RuleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ruleID := r.PathValue("id")

		if ruleID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}
		if _, err := uuid.Parse(ruleID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid id (uuid)")
			return
		}
		if err := svc.DeleteRule(ctx, ruleID); err != nil {
			switch {
			case errors.Is(err, domain.ErrNotFound):
				json.HandleError(w, http.StatusNotFound, err, "rule not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "failed to delete rule")
			}
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
