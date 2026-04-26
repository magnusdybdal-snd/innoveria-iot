package handlers_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers"
)

type mockRuleService struct {
	t              *testing.T
	createRuleFunc func(ctx context.Context, rule domain.AggregationRule) (string, error)
	getRulesFunc   func(ctx context.Context, companyID string) ([]domain.AggregationRule, error)
	deleteRuleFunc func(ctx context.Context, ruleID string) error
}

func (m *mockRuleService) CreateRule(ctx context.Context, rule domain.AggregationRule) (string, error) {
	if m.createRuleFunc == nil {
		m.t.Fatal("unexpected call to CreateRule")
	}
	return m.createRuleFunc(ctx, rule)
}

func (m *mockRuleService) GetRules(ctx context.Context, companyID string) ([]domain.AggregationRule, error) {
	if m.getRulesFunc == nil {
		m.t.Fatal("unexpected call to GetRules")
	}
	return m.getRulesFunc(ctx, companyID)
}

func (m *mockRuleService) DeleteRule(ctx context.Context, ruleID string) error {
	if m.deleteRuleFunc == nil {
		m.t.Fatal("unexpected call to DeleteRule")
	}
	return m.deleteRuleFunc(ctx, ruleID)
}

const validRuleBody = `{
	"name": "Test Rule",
	"context_type": "energy",
	"measurement_type": "watt",
	"aggregation_method": "AVG",
	"time_bucket_minutes": 15,
	"is_active": true
}`

// TestCreateRule_ValidBody_Returns201 verifies that a valid request body returns 201 Created.
func TestCreateRule_ValidBody_Returns201(t *testing.T) {
	svc := &mockRuleService{
		t: t,
		createRuleFunc: func(_ context.Context, _ domain.AggregationRule) (string, error) {
			return "some-uuid", nil
		},
	}

	req := withAuthHeaders(httptest.NewRequest(http.MethodPost, "/api/v1/context/rules", strings.NewReader(validRuleBody)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.CreateRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rec.Code)
	}
}

// TestCreateRule_MalformedJSON_Returns400 verifies that malformed JSON returns 400 Bad Request.
func TestCreateRule_MalformedJSON_Returns400(t *testing.T) {
	svc := &mockRuleService{t: t}

	req := withAuthHeaders(httptest.NewRequest(http.MethodPost, "/api/v1/context/rules", strings.NewReader(`{not valid json`)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.CreateRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestCreateRule_EmptyBody_Returns400 verifies that an empty body returns 400 Bad Request.
func TestCreateRule_EmptyBody_Returns400(t *testing.T) {
	svc := &mockRuleService{t: t}

	req := withAuthHeaders(httptest.NewRequest(http.MethodPost, "/api/v1/context/rules", strings.NewReader("")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.CreateRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestCreateRule_MissingRequiredFields_Returns400 verifies that missing required string fields returns 400 Bad Request.
func TestCreateRule_MissingRequiredFields_Returns400(t *testing.T) {
	svc := &mockRuleService{t: t}

	cases := []struct {
		name string
		body string
	}{
		{"missing name", `{"context_type":"energy","measurement_type":"watt","aggregation_method":"AVG","time_bucket_minutes":15}`},
		{"missing context_type", `{"name":"Test","measurement_type":"watt","aggregation_method":"AVG","time_bucket_minutes":15}`},
		{"missing measurement_type", `{"name":"Test","context_type":"energy","aggregation_method":"AVG","time_bucket_minutes":15}`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := withAuthHeaders(httptest.NewRequest(http.MethodPost, "/api/v1/context/rules", strings.NewReader(tc.body)))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handlers.CreateRule(svc).ServeHTTP(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Errorf("expected 400, got %d", rec.Code)
			}
		})
	}
}

// TestCreateRule_InvalidAggregationMethod_Returns400 verifies that an invalid aggregation_method returns 400 Bad Request.
func TestCreateRule_InvalidAggregationMethod_Returns400(t *testing.T) {
	svc := &mockRuleService{t: t}

	body := `{
		"name": "Test Rule",
		"context_type": "energy",
		"measurement_type": "watt",
		"aggregation_method": "INVALID",
		"time_bucket_minutes": 15,
		"is_active": true
	}`

	req := withAuthHeaders(httptest.NewRequest(http.MethodPost, "/api/v1/context/rules", strings.NewReader(body)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.CreateRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestCreateRule_ZeroTimeBucket_Returns400 verifies that a time_bucket_minutes of 0 returns 400 Bad Request.
func TestCreateRule_ZeroTimeBucket_Returns400(t *testing.T) {
	svc := &mockRuleService{t: t}

	body := `{
		"name": "Test Rule",
		"context_type": "energy",
		"measurement_type": "watt",
		"aggregation_method": "AVG",
		"time_bucket_minutes": 0,
		"is_active": true
	}`

	req := withAuthHeaders(httptest.NewRequest(http.MethodPost, "/api/v1/context/rules", strings.NewReader(body)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.CreateRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestCreateRule_NegativeTimeBucket_Returns400 verifies that a negative time_bucket_minutes returns 400 Bad Request.
func TestCreateRule_NegativeTimeBucket_Returns400(t *testing.T) {
	svc := &mockRuleService{t: t}

	body := `{
		"name": "Test Rule",
		"context_type": "energy",
		"measurement_type": "watt",
		"aggregation_method": "AVG",
		"time_bucket_minutes": -10,
		"is_active": true
	}`

	req := withAuthHeaders(httptest.NewRequest(http.MethodPost, "/api/v1/context/rules", strings.NewReader(body)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.CreateRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestCreateRule_DuplicateRule_Returns409 verifies that a duplicate (company_id, context_type) returns 409 Conflict.
func TestCreateRule_DuplicateRule_Returns409(t *testing.T) {
	svc := &mockRuleService{
		t: t,
		createRuleFunc: func(_ context.Context, _ domain.AggregationRule) (string, error) {
			return "", domain.ErrConflict
		},
	}

	req := withAuthHeaders(httptest.NewRequest(http.MethodPost, "/api/v1/context/rules", strings.NewReader(validRuleBody)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.CreateRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusConflict {
		t.Errorf("expected 409, got %d", rec.Code)
	}
}

// TestCreateRule_ServiceError_Returns500 verifies that a service error returns 500 Internal Server Error.
func TestCreateRule_ServiceError_Returns500(t *testing.T) {
	svc := &mockRuleService{
		t: t,
		createRuleFunc: func(_ context.Context, _ domain.AggregationRule) (string, error) {
			return "", errors.New("db error")
		},
	}

	req := withAuthHeaders(httptest.NewRequest(http.MethodPost, "/api/v1/context/rules", strings.NewReader(validRuleBody)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.CreateRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}

// Tests of DELETE handler for aggregation rules.

const validRuleID = "b1111111-0000-0000-0000-000000000001"

// TestDeleteRule_Success_Returns204 verifies that a valid rule ID returns 204 No Content.
func TestDeleteRule_Success_Returns204(t *testing.T) {
	svc := &mockRuleService{
		t: t,
		deleteRuleFunc: func(_ context.Context, _ string) error {
			return nil
		},
	}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/context/rules/"+validRuleID, nil)
	req.SetPathValue("id", validRuleID)
	rec := httptest.NewRecorder()

	handlers.DeleteRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", rec.Code)
	}
}

// TestDeleteRule_MissingID_Returns400 verifies that a missing rule ID path value returns 400 Bad Request.
func TestDeleteRule_MissingID_Returns400(t *testing.T) {
	svc := &mockRuleService{t: t}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/context/rules/", nil)
	rec := httptest.NewRecorder()

	handlers.DeleteRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestDeleteRule_InvalidUUID_Returns400 verifies that a non-UUID rule ID returns 400 Bad Request.
func TestDeleteRule_InvalidUUID_Returns400(t *testing.T) {
	svc := &mockRuleService{t: t}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/context/rules/not-a-uuid", nil)
	req.SetPathValue("id", "not-a-uuid")
	rec := httptest.NewRecorder()

	handlers.DeleteRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestDeleteRule_NotFound_Returns404 verifies that a non-existent rule ID returns 404 Not Found.
func TestDeleteRule_NotFound_Returns404(t *testing.T) {
	svc := &mockRuleService{
		t: t,
		deleteRuleFunc: func(_ context.Context, _ string) error {
			return domain.ErrNotFound
		},
	}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/context/rules/"+validRuleID, nil)
	req.SetPathValue("id", validRuleID)
	rec := httptest.NewRecorder()

	handlers.DeleteRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", rec.Code)
	}
}

// TestDeleteRule_ServiceError_Returns500 verifies that an unexpected service error returns 500 Internal Server Error.
func TestDeleteRule_ServiceError_Returns500(t *testing.T) {
	svc := &mockRuleService{
		t: t,
		deleteRuleFunc: func(_ context.Context, _ string) error {
			return errors.New("db error")
		},
	}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/context/rules/"+validRuleID, nil)
	req.SetPathValue("id", validRuleID)
	rec := httptest.NewRecorder()

	handlers.DeleteRule(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}
