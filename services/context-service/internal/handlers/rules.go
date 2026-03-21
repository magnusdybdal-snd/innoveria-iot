package handlers

import (
	"fmt"
	"net/http"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// GetRules retrieves all aggregation rules for a given company ID.
// @Summary 		Get Aggregation Rules
// @Tags 			Rules
// @Accept 			json
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
			json.HandleError(w, http.StatusInternalServerError, err, "failed to retrieve rules")
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
				AggregationMethod: rule.AggregationMethod,
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
