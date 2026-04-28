package handlers_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers"
)

// mockSensorService is a test double for domain.SensorService.
// Each method delegates to a function field; an unset field causes the test to fail immediately,
// catching unexpected calls to the service.
type mockSensorService struct {
	t                           *testing.T
	getAllFunc                  func(ctx context.Context, companyID string) ([]domain.Sensor, error)
	getByIDFunc                 func(ctx context.Context, companyID string, sensorID string) (domain.Sensor, error)
	getByProductionResourceFunc func(ctx context.Context, companyID string, productionResourceID int64) ([]domain.Sensor, error)
	createFunc                  func(ctx context.Context, payload domain.Sensor) error
	updateFunc                  func(ctx context.Context, companyID string, deviceID string, payload domain.Sensor) error
	deleteFunc                  func(ctx context.Context, companyID string, deviceID string) error
	getSampleEUIFunc            func(ctx context.Context, chirpstackProfileID string) (string, error)
}

func (m *mockSensorService) GetAll(ctx context.Context, companyID string) ([]domain.Sensor, error) {
	if m.getAllFunc == nil {
		m.t.Fatal("unexpected call to GetAll")
	}
	return m.getAllFunc(ctx, companyID)
}

func (m *mockSensorService) GetByID(ctx context.Context, companyID string, sensorID string) (domain.Sensor, error) {
	if m.getByIDFunc == nil {
		m.t.Fatal("unexpected call to GetByID")
	}
	return m.getByIDFunc(ctx, companyID, sensorID)
}

func (m *mockSensorService) GetByProductionResourceID(ctx context.Context, companyID string, productionResourceID int64) ([]domain.Sensor, error) {
	if m.getByProductionResourceFunc == nil {
		m.t.Fatal("unexpected call to GetByProductionResourceID")
	}
	return m.getByProductionResourceFunc(ctx, companyID, productionResourceID)
}

func (m *mockSensorService) Create(ctx context.Context, payload domain.Sensor) error {
	if m.createFunc == nil {
		m.t.Fatal("unexpected call to Create")
	}
	return m.createFunc(ctx, payload)
}

func (m *mockSensorService) Update(ctx context.Context, companyID string, deviceID string, payload domain.Sensor) error {
	if m.updateFunc == nil {
		m.t.Fatal("unexpected call to Update")
	}
	return m.updateFunc(ctx, companyID, deviceID, payload)
}

func (m *mockSensorService) Delete(ctx context.Context, companyID string, deviceID string) error {
	if m.deleteFunc == nil {
		m.t.Fatal("unexpected call to Delete")
	}
	return m.deleteFunc(ctx, companyID, deviceID)
}

func (m *mockSensorService) GetSampleEUI(ctx context.Context, chirpstackProfileID string) (string, error) {
	if m.getSampleEUIFunc == nil {
		m.t.Fatal("unexpected call to GetSampleEUI")
	}
	return m.getSampleEUIFunc(ctx, chirpstackProfileID)
}

const validSensorBody = `{
	"company_id":       "a0000000-0000-0000-0000-000000000001",
	"name":             "Test Sensor",
	"device_eui":       "b000000000000001",
	"app_key":          "00000000000000000000000000000001",
	"device_profile_id":"f0000000-0000-0000-0000-000000000001",
	"factory_id":       "f1000000-0000-0000-0000-000000000001",
	"factory_area_id":  "a1000000-0000-0000-0000-000000000001"
}`

const validSensorID = "c0000000-0000-0000-0000-000000000001"

// ── GetSensors ────────────────────────────────────────────────────────────────

// TestGetSensors_NoFilter_Returns200 verifies that a request with no filter calls GetAll and returns 200.
func TestGetSensors_NoFilter_Returns200(t *testing.T) {
	svc := &mockSensorService{
		t: t,
		getAllFunc: func(_ context.Context, _ string) ([]domain.Sensor, error) {
			return []domain.Sensor{}, nil
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sensors", nil)
	rec := httptest.NewRecorder()

	handlers.GetSensors(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}

// TestGetSensors_ValidProductionResourceID_Returns200 verifies that a numeric production_resource_id
// calls GetByProductionResourceID and returns 200.
func TestGetSensors_ValidProductionResourceID_Returns200(t *testing.T) {
	svc := &mockSensorService{
		t: t,
		getByProductionResourceFunc: func(_ context.Context, _ string, id int64) ([]domain.Sensor, error) {
			if id != 1 {
				t.Errorf("expected production resource ID 1, got %d", id)
			}
			return []domain.Sensor{}, nil
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sensors?production_resource_id=1", nil)
	rec := httptest.NewRecorder()

	handlers.GetSensors(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}

// TestGetSensors_NonNumericProductionResourceID_Returns400 verifies that a non-numeric
// production_resource_id returns 400 Bad Request without calling the service.
func TestGetSensors_NonNumericProductionResourceID_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sensors?production_resource_id=not-a-number", nil)
	rec := httptest.NewRecorder()

	handlers.GetSensors(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestGetSensors_UUIDProductionResourceID_Returns400 verifies that the old UUID format is now
// rejected with 400 Bad Request (production_resource_id is int64, not uuid).
func TestGetSensors_UUIDProductionResourceID_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sensors?production_resource_id=a0000000-0000-0000-0000-000000000001", nil)
	rec := httptest.NewRecorder()

	handlers.GetSensors(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestGetSensors_NonPositiveProductionResourceID_Returns400 verifies that zero and negative values
// are rejected — ERP production resource IDs start from 1.
func TestGetSensors_NonPositiveProductionResourceID_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	for _, id := range []string{"0", "-1", "-100"} {
		t.Run("id="+id, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/sensors?production_resource_id="+id, nil)
			rec := httptest.NewRecorder()

			handlers.GetSensors(svc).ServeHTTP(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Errorf("expected 400, got %d", rec.Code)
			}
		})
	}
}

// TestGetSensors_ServiceError_Returns500 verifies that a service error returns 500.
func TestGetSensors_ServiceError_Returns500(t *testing.T) {
	svc := &mockSensorService{
		t: t,
		getAllFunc: func(_ context.Context, _ string) ([]domain.Sensor, error) {
			return nil, errors.New("db error")
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sensors", nil)
	rec := httptest.NewRecorder()

	handlers.GetSensors(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}

// ── PostSensor ────────────────────────────────────────────────────────────────

// TestPostSensor_ValidBody_Returns201 verifies that a complete, valid request body returns 201 Created.
func TestPostSensor_ValidBody_Returns201(t *testing.T) {
	svc := &mockSensorService{
		t: t,
		createFunc: func(_ context.Context, _ domain.Sensor) error {
			return nil
		},
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/sensors", strings.NewReader(validSensorBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.PostSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rec.Code)
	}
}

// TestPostSensor_MalformedJSON_Returns400 verifies that malformed JSON returns 400 Bad Request.
func TestPostSensor_MalformedJSON_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/sensors", strings.NewReader(`{not valid json`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.PostSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestPostSensor_MissingRequiredFields_Returns400 verifies that omitting any required field returns 400.
func TestPostSensor_MissingRequiredFields_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	cases := []struct {
		name string
		body string
	}{
		{"missing company_id", `{"name":"S","device_eui":"b000000000000001","app_key":"00000000000000000000000000000001","device_profile_id":"f0000000-0000-0000-0000-000000000001","factory_id":"f1000000-0000-0000-0000-000000000001","factory_area_id":"a1000000-0000-0000-0000-000000000001"}`},
		{"missing name", `{"company_id":"a0000000-0000-0000-0000-000000000001","device_eui":"b000000000000001","app_key":"00000000000000000000000000000001","device_profile_id":"f0000000-0000-0000-0000-000000000001","factory_id":"f1000000-0000-0000-0000-000000000001","factory_area_id":"a1000000-0000-0000-0000-000000000001"}`},
		{"missing device_eui", `{"company_id":"a0000000-0000-0000-0000-000000000001","name":"S","app_key":"00000000000000000000000000000001","device_profile_id":"f0000000-0000-0000-0000-000000000001","factory_id":"f1000000-0000-0000-0000-000000000001","factory_area_id":"a1000000-0000-0000-0000-000000000001"}`},
		{"missing app_key", `{"company_id":"a0000000-0000-0000-0000-000000000001","name":"S","device_eui":"b000000000000001","device_profile_id":"f0000000-0000-0000-0000-000000000001","factory_id":"f1000000-0000-0000-0000-000000000001","factory_area_id":"a1000000-0000-0000-0000-000000000001"}`},
		{"missing device_profile_id", `{"company_id":"a0000000-0000-0000-0000-000000000001","name":"S","device_eui":"b000000000000001","app_key":"00000000000000000000000000000001","factory_id":"f1000000-0000-0000-0000-000000000001","factory_area_id":"a1000000-0000-0000-0000-000000000001"}`},
		{"missing factory_id", `{"company_id":"a0000000-0000-0000-0000-000000000001","name":"S","device_eui":"b000000000000001","app_key":"00000000000000000000000000000001","device_profile_id":"f0000000-0000-0000-0000-000000000001","factory_area_id":"a1000000-0000-0000-0000-000000000001"}`},
		{"missing factory_area_id", `{"company_id":"a0000000-0000-0000-0000-000000000001","name":"S","device_eui":"b000000000000001","app_key":"00000000000000000000000000000001","device_profile_id":"f0000000-0000-0000-0000-000000000001","factory_id":"f1000000-0000-0000-0000-000000000001"}`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/sensors", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handlers.PostSensor(svc).ServeHTTP(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Errorf("expected 400, got %d", rec.Code)
			}
		})
	}
}

// TestPostSensor_ServiceError_Returns500 verifies that a service-level error returns 500.
func TestPostSensor_ServiceError_Returns500(t *testing.T) {
	svc := &mockSensorService{
		t: t,
		createFunc: func(_ context.Context, _ domain.Sensor) error {
			return errors.New("db error")
		},
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/sensors", strings.NewReader(validSensorBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.PostSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}

// ── PatchSensor ───────────────────────────────────────────────────────────────

// TestPatchSensor_ValidBody_Returns204 verifies that a valid sensor ID and update payload returns 204.
func TestPatchSensor_ValidBody_Returns204(t *testing.T) {
	svc := &mockSensorService{
		t: t,
		updateFunc: func(_ context.Context, _, _ string, _ domain.Sensor) error {
			return nil
		},
	}

	body := `{"name":"Updated Name"}`
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/sensors/"+validSensorID, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", validSensorID)
	rec := httptest.NewRecorder()

	handlers.PatchSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", rec.Code)
	}
}

// TestPatchSensor_MissingID_Returns400 verifies that an empty path value returns 400.
func TestPatchSensor_MissingID_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	req := httptest.NewRequest(http.MethodPatch, "/api/v1/sensors/", strings.NewReader(`{"name":"x"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handlers.PatchSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestPatchSensor_InvalidUUID_Returns400 verifies that a non-UUID sensor ID returns 400.
func TestPatchSensor_InvalidUUID_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	req := httptest.NewRequest(http.MethodPatch, "/api/v1/sensors/not-a-uuid", strings.NewReader(`{"name":"x"}`))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", "not-a-uuid")
	rec := httptest.NewRecorder()

	handlers.PatchSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestPatchSensor_NoFields_Returns400 verifies that a body with all nil fields returns 400.
func TestPatchSensor_NoFields_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	req := httptest.NewRequest(http.MethodPatch, "/api/v1/sensors/"+validSensorID, strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", validSensorID)
	rec := httptest.NewRecorder()

	handlers.PatchSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestPatchSensor_SetsProductionResource_PassesValueToService verifies that the production_resource
// field is forwarded to the service so it reaches the UPDATE query.
// This guards against the previous bug where production_resource_id was silently dropped from the SET clause.
func TestPatchSensor_SetsProductionResource_PassesValueToService(t *testing.T) {
	var gotProductionResource *int64
	svc := &mockSensorService{
		t: t,
		updateFunc: func(_ context.Context, _, _ string, payload domain.Sensor) error {
			gotProductionResource = payload.ProductionResource
			return nil
		},
	}

	body := `{"production_resource": 42}`
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/sensors/"+validSensorID, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", validSensorID)
	rec := httptest.NewRecorder()

	handlers.PatchSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", rec.Code)
	}
	if gotProductionResource == nil || *gotProductionResource != 42 {
		t.Errorf("expected ProductionResource=42, got %v", gotProductionResource)
	}
}

// TestPatchSensor_ClearsProductionResource_PassesZeroToService verifies that sending
// production_resource: 0 forwards the sentinel value to the service, which interprets it as "clear".
func TestPatchSensor_ClearsProductionResource_PassesZeroToService(t *testing.T) {
	var gotProductionResource *int64
	svc := &mockSensorService{
		t: t,
		updateFunc: func(_ context.Context, _, _ string, payload domain.Sensor) error {
			gotProductionResource = payload.ProductionResource
			return nil
		},
	}

	body := `{"production_resource": 0}`
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/sensors/"+validSensorID, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", validSensorID)
	rec := httptest.NewRecorder()

	handlers.PatchSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", rec.Code)
	}
	if gotProductionResource == nil || *gotProductionResource != 0 {
		t.Errorf("expected ProductionResource=0 (clear sentinel), got %v", gotProductionResource)
	}
}

// TestPatchSensor_NegativeProductionResource_Returns400 verifies that a negative production_resource
// is rejected — only 0 (clear) and positive integers are valid.
func TestPatchSensor_NegativeProductionResource_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	body := `{"production_resource": -1}`
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/sensors/"+validSensorID, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", validSensorID)
	rec := httptest.NewRecorder()

	handlers.PatchSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestPatchSensor_ServiceError_Returns500 verifies that a service error returns 500.
func TestPatchSensor_ServiceError_Returns500(t *testing.T) {
	svc := &mockSensorService{
		t: t,
		updateFunc: func(_ context.Context, _, _ string, _ domain.Sensor) error {
			return errors.New("db error")
		},
	}

	body := `{"name":"Updated Name"}`
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/sensors/"+validSensorID, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", validSensorID)
	rec := httptest.NewRecorder()

	handlers.PatchSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}

// ── DeleteSensor ──────────────────────────────────────────────────────────────

// TestDeleteSensor_ValidID_Returns204 verifies that a valid sensor UUID returns 204.
func TestDeleteSensor_ValidID_Returns204(t *testing.T) {
	svc := &mockSensorService{
		t: t,
		deleteFunc: func(_ context.Context, _, _ string) error {
			return nil
		},
	}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/sensors/"+validSensorID, nil)
	req.SetPathValue("id", validSensorID)
	rec := httptest.NewRecorder()

	handlers.DeleteSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", rec.Code)
	}
}

// TestDeleteSensor_MissingID_Returns400 verifies that an empty path value returns 400.
func TestDeleteSensor_MissingID_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/sensors/", nil)
	rec := httptest.NewRecorder()

	handlers.DeleteSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestDeleteSensor_InvalidUUID_Returns400 verifies that a non-UUID sensor ID returns 400.
func TestDeleteSensor_InvalidUUID_Returns400(t *testing.T) {
	svc := &mockSensorService{t: t}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/sensors/not-a-uuid", nil)
	req.SetPathValue("id", "not-a-uuid")
	rec := httptest.NewRecorder()

	handlers.DeleteSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// TestDeleteSensor_ServiceError_Returns500 verifies that a service error returns 500.
func TestDeleteSensor_ServiceError_Returns500(t *testing.T) {
	svc := &mockSensorService{
		t: t,
		deleteFunc: func(_ context.Context, _, _ string) error {
			return errors.New("db error")
		},
	}

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/sensors/"+validSensorID, nil)
	req.SetPathValue("id", validSensorID)
	rec := httptest.NewRecorder()

	handlers.DeleteSensor(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}
