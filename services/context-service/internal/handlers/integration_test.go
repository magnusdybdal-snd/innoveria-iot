//go:build integration

package handlers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"

	"innoveria-iot/context-service/internal/clients"
	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/server"
	"innoveria-iot/context-service/internal/services"
)

// inMemoryRuleRepo is a thread-safe in-memory implementation of domain.RuleRepository for tests.
type inMemoryRuleRepo struct {
	mu    sync.Mutex
	rules map[string]domain.AggregationRule
}

// newInMemoryRuleRepo creates an empty inMemoryRuleRepo.
func newInMemoryRuleRepo() *inMemoryRuleRepo {
	return &inMemoryRuleRepo{rules: make(map[string]domain.AggregationRule)}
}

// GetByCompanyID returns all rules that belong to the given company.
func (r *inMemoryRuleRepo) GetByCompanyID(_ context.Context, companyID string) ([]domain.AggregationRule, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var out []domain.AggregationRule
	for _, rule := range r.rules {
		if rule.CompanyID == companyID {
			out = append(out, rule)
		}
	}
	if out == nil {
		out = []domain.AggregationRule{}
	}
	return out, nil
}

// GetByID returns the rule with the given ID, or domain.ErrNotFound.
func (r *inMemoryRuleRepo) GetByID(_ context.Context, ruleID string) (domain.AggregationRule, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	rule, ok := r.rules[ruleID]
	if !ok {
		return domain.AggregationRule{}, domain.ErrNotFound
	}
	return rule, nil
}

// Create assigns a new UUID to the rule, stores it, and returns the ID.
func (r *inMemoryRuleRepo) Create(_ context.Context, rule domain.AggregationRule) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	rule.ID = uuid.New().String()
	r.rules[rule.ID] = rule
	return rule.ID, nil
}

// DeleteByID removes the rule, returning domain.ErrNotFound if it does not exist.
func (r *inMemoryRuleRepo) DeleteByID(_ context.Context, ruleID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.rules[ruleID]; !ok {
		return domain.ErrNotFound
	}
	delete(r.rules, ruleID)
	return nil
}

// testServerResult bundles the main test server with the shared rule repo so tests can pre-populate it.
type testServerResult struct {
	srv      *httptest.Server
	ruleRepo *inMemoryRuleRepo
}

// newTestServer wires up three downstream stubs, real clients, real services, and the router,
// then returns the running main httptest.Server together with the rule repo.
func newTestServer(t *testing.T, erpHandler, deviceHandler, collectionHandler http.Handler) testServerResult {
	t.Helper()

	erpStub := httptest.NewServer(erpHandler)
	deviceStub := httptest.NewServer(deviceHandler)
	collectionStub := httptest.NewServer(collectionHandler)

	repo := newInMemoryRuleRepo()

	erpClient := clients.NewERPClient(erpStub.URL)
	deviceClient := clients.NewDeviceClient(deviceStub.URL, deviceStub.URL)
	collectionClient := clients.NewCollectionClient(collectionStub.URL)

	contextSvc := services.NewContextServiceImpl(collectionClient, erpClient, deviceClient, repo)
	ruleSvc := services.NewRuleServiceImpl(repo)

	mux := server.NewRouter(contextSvc, ruleSvc)
	mainSrv := httptest.NewServer(mux)

	t.Cleanup(func() {
		mainSrv.Close()
		erpStub.Close()
		deviceStub.Close()
		collectionStub.Close()
	})

	return testServerResult{srv: mainSrv, ruleRepo: repo}
}

// addAuthHeaders copies the company-A auth headers onto the request.
func addAuthHeaders(r *http.Request) *http.Request {
	r.Header.Set("X-Auth-User-Id", "u0000000-0000-0000-0000-000000000001")
	r.Header.Set("X-Auth-Company-Id", "a0000000-0000-0000-0000-000000000001")
	r.Header.Set("X-Auth-Role", "FACTORY_WORKER")
	return r
}

// addCompanyBHeaders sets auth headers scoped to company B.
func addCompanyBHeaders(r *http.Request) *http.Request {
	r.Header.Set("X-Auth-User-Id", "u0000000-0000-0000-0000-000000000002")
	r.Header.Set("X-Auth-Company-Id", "b0000000-0000-0000-0000-000000000001")
	r.Header.Set("X-Auth-Role", "FACTORY_WORKER")
	return r
}

// TestIntegration_OrderContext_DegradedOperation verifies that when the device service
// returns 5xx for one operation's production resource, that operation is Degraded: true
// while the sibling operation with a healthy device response is Degraded: false.
func TestIntegration_OrderContext_DegradedOperation(t *testing.T) {
	now := time.Now()
	actualStart := now.Add(-1 * time.Hour)

	erpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/erp/orders/42" {
			http.NotFound(w, r)
			return
		}
		resp := map[string]any{
			"order": map[string]any{
				"id":                  42,
				"company_id":          "a0000000-0000-0000-0000-000000000001",
				"order_number":        "MO-42",
				"part_id":             "P1",
				"part_description":    "Part One",
				"planned_start_date":  now.Add(-2 * time.Hour).Format(time.RFC3339),
				"planned_finish_date": now.Add(2 * time.Hour).Format(time.RFC3339),
				"actual_start_date":   actualStart.Format(time.RFC3339),
				"actual_finish_date":  nil,
				"status":              "started",
				"priority":            1,
				"received_at":         now.Format(time.RFC3339),
			},
			"operations": []map[string]any{
				{
					"id":                         1,
					"production_resource_id":     1,
					"order_id":                   42,
					"planned_start_date":         now.Add(-2 * time.Hour).Format(time.RFC3339),
					"planned_finish_date":        now.Add(2 * time.Hour).Format(time.RFC3339),
					"actual_start_date":          actualStart.Format(time.RFC3339),
					"actual_finish_date":         nil,
					"status":                     "started",
					"production_resource_status": "started",
					"received_at":                now.Format(time.RFC3339),
					"reports":                    []any{},
				},
				{
					"id":                         2,
					"production_resource_id":     2,
					"order_id":                   42,
					"planned_start_date":         now.Add(-2 * time.Hour).Format(time.RFC3339),
					"planned_finish_date":        now.Add(2 * time.Hour).Format(time.RFC3339),
					"actual_start_date":          actualStart.Format(time.RFC3339),
					"actual_finish_date":         nil,
					"status":                     "started",
					"production_resource_status": "started",
					"received_at":                now.Format(time.RFC3339),
					"reports":                    []any{},
				},
			},
			"production_resources": []map[string]any{
				{"id": 1, "number": "WC-1", "description": "Work Center 1", "type": "machine", "received_at": now.Format(time.RFC3339)},
				{"id": 2, "number": "WC-2", "description": "Work Center 2", "type": "machine", "received_at": now.Format(time.RFC3339)},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	deviceHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/metrics") {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"total_count": 0, "metrics": []any{}})
			return
		}
		prodResourceID := r.URL.Query().Get("production_resource_id")
		switch prodResourceID {
		case "1":
			resp := map[string]any{
				"total_count": 1,
				"sensors": []map[string]any{
					{
						"id":                  "sensor-aaa",
						"name":                "Sensor A",
						"device_eui":          "AA:BB:CC:DD:EE:FF:00:01",
						"production_resource": 1,
						"electricity_sensor":  false,
						"voltage":             nil,
					},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(resp)
		case "2":
			http.Error(w, "internal server error", http.StatusInternalServerError)
		default:
			http.NotFound(w, r)
		}
	})

	collectionHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode([]any{})
	})

	ts := newTestServer(t, erpHandler, deviceHandler, collectionHandler)

	req, err := http.NewRequest(http.MethodGet, ts.srv.URL+"/api/v1/context/orders/42/context", nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	addAuthHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var body struct {
		Operations []struct {
			Operation struct {
				ProductionResource struct {
					ID int64 `json:"id"`
				} `json:"production_resource"`
			} `json:"operation"`
			Sensors  []json.RawMessage `json:"sensors"`
			Degraded bool              `json:"degraded"`
		} `json:"operations"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(body.Operations) != 2 {
		t.Fatalf("expected 2 operations, got %d", len(body.Operations))
	}

	for _, op := range body.Operations {
		switch op.Operation.ProductionResource.ID {
		case 1:
			if op.Degraded {
				t.Errorf("operation with production_resource_id=1 should not be degraded")
			}
		case 2:
			if !op.Degraded {
				t.Errorf("operation with production_resource_id=2 should be degraded")
			}
		default:
			t.Errorf("unexpected production_resource_id: %d", op.Operation.ProductionResource.ID)
		}
	}
}

// TestIntegration_OrderContext_EmptySensors_NotDegraded verifies that a 404 from
// the device service results in an empty sensors array and Degraded: false.
func TestIntegration_OrderContext_EmptySensors_NotDegraded(t *testing.T) {
	now := time.Now()
	actualStart := now.Add(-1 * time.Hour)

	erpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/erp/orders/43" {
			http.NotFound(w, r)
			return
		}
		resp := map[string]any{
			"order": map[string]any{
				"id":                  43,
				"company_id":          "a0000000-0000-0000-0000-000000000001",
				"order_number":        "MO-43",
				"part_id":             "P2",
				"part_description":    "Part Two",
				"planned_start_date":  now.Add(-2 * time.Hour).Format(time.RFC3339),
				"planned_finish_date": now.Add(2 * time.Hour).Format(time.RFC3339),
				"actual_start_date":   actualStart.Format(time.RFC3339),
				"actual_finish_date":  nil,
				"status":              "started",
				"priority":            1,
				"received_at":         now.Format(time.RFC3339),
			},
			"operations": []map[string]any{
				{
					"id":                         10,
					"production_resource_id":     5,
					"order_id":                   43,
					"planned_start_date":         now.Add(-2 * time.Hour).Format(time.RFC3339),
					"planned_finish_date":        now.Add(2 * time.Hour).Format(time.RFC3339),
					"actual_start_date":          actualStart.Format(time.RFC3339),
					"actual_finish_date":         nil,
					"status":                     "started",
					"production_resource_status": "started",
					"received_at":                now.Format(time.RFC3339),
					"reports":                    []any{},
				},
			},
			"production_resources": []map[string]any{
				{"id": 5, "number": "WC-5", "description": "Work Center 5", "type": "machine", "received_at": now.Format(time.RFC3339)},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	deviceHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	collectionHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode([]any{})
	})

	ts := newTestServer(t, erpHandler, deviceHandler, collectionHandler)

	req, err := http.NewRequest(http.MethodGet, ts.srv.URL+"/api/v1/context/orders/43/context", nil)
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}
	addAuthHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var body struct {
		Operations []struct {
			Sensors  []json.RawMessage `json:"sensors"`
			Degraded bool              `json:"degraded"`
		} `json:"operations"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(body.Operations) != 1 {
		t.Fatalf("expected 1 operation, got %d", len(body.Operations))
	}

	op := body.Operations[0]
	if op.Degraded {
		t.Errorf("operation should not be degraded when device service returns 404")
	}
	if len(op.Sensors) != 0 {
		t.Errorf("expected 0 sensors, got %d", len(op.Sensors))
	}
}

// TestIntegration_Rules_CrossCompanyIsolation verifies that rules created by company A
// are invisible to company B.
func TestIntegration_Rules_CrossCompanyIsolation(t *testing.T) {
	ts := newTestServer(t, http.NotFoundHandler(), http.NotFoundHandler(), http.NotFoundHandler())

	ruleBody := `{"name":"Test Rule","context_type":"energy","measurement_type":"electric_current","aggregation_method":"AVG","time_bucket_minutes":15,"is_active":true}`

	// Step 1: create rule as company A.
	createReq, err := http.NewRequest(http.MethodPost, ts.srv.URL+"/api/v1/context/rules", strings.NewReader(ruleBody))
	if err != nil {
		t.Fatalf("failed to build create request: %v", err)
	}
	createReq.Header.Set("Content-Type", "application/json")
	addAuthHeaders(createReq)

	createResp, err := http.DefaultClient.Do(createReq)
	if err != nil {
		t.Fatalf("create request failed: %v", err)
	}
	defer createResp.Body.Close()

	if createResp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201 on create, got %d", createResp.StatusCode)
	}

	var createdID string
	if err := json.NewDecoder(createResp.Body).Decode(&createdID); err != nil {
		t.Fatalf("failed to decode created rule ID: %v", err)
	}
	if createdID == "" {
		t.Fatal("expected non-empty rule ID from create")
	}

	// Step 2: list rules as company A — the rule must be present.
	listAReq, err := http.NewRequest(http.MethodGet, ts.srv.URL+"/api/v1/context/rules", nil)
	if err != nil {
		t.Fatalf("failed to build list-A request: %v", err)
	}
	addAuthHeaders(listAReq)

	listAResp, err := http.DefaultClient.Do(listAReq)
	if err != nil {
		t.Fatalf("list-A request failed: %v", err)
	}
	defer listAResp.Body.Close()

	if listAResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 on list for company A, got %d", listAResp.StatusCode)
	}

	var rulesA []struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(listAResp.Body).Decode(&rulesA); err != nil {
		t.Fatalf("failed to decode company A rules: %v", err)
	}

	foundInA := false
	for _, r := range rulesA {
		if r.ID == createdID {
			foundInA = true
			break
		}
	}
	if !foundInA {
		t.Errorf("rule %s not found in company A list", createdID)
	}

	// Step 3: list rules as company B — the rule must not appear.
	listBReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/context/rules", ts.srv.URL), nil)
	if err != nil {
		t.Fatalf("failed to build list-B request: %v", err)
	}
	addCompanyBHeaders(listBReq)

	listBResp, err := http.DefaultClient.Do(listBReq)
	if err != nil {
		t.Fatalf("list-B request failed: %v", err)
	}
	defer listBResp.Body.Close()

	if listBResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 on list for company B, got %d", listBResp.StatusCode)
	}

	var rulesB []struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(listBResp.Body).Decode(&rulesB); err != nil {
		t.Fatalf("failed to decode company B rules: %v", err)
	}

	for _, r := range rulesB {
		if r.ID == createdID {
			t.Errorf("rule %s from company A leaked into company B list", createdID)
		}
	}
}
