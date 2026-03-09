// Package handlers provides HTTP handlers for the API gateway.
package handlers

import (
	"context"
	"innoveria-iot/pkg/httpclient"
	"innoveria-iot/pkg/json"
	"maps"
	"net/http"
)

// swaggerSpec represents the structure of a Swagger 2.0 JSON spec.
// Each service generates one of these via swag init.
// We only include the fields we need for merging — the rest are ignored.
type swaggerSpec struct {
	Swagger     string         `json:"swagger"`
	Info        map[string]any `json:"info"`
	Host        string         `json:"host"`
	BasePath    string         `json:"basePath"`    // e.g. "/api/v1/device"
	Paths       map[string]any `json:"paths"`       // e.g. {"/sensors": {...}}
	Definitions map[string]any `json:"definitions"` // e.g. {"dto.SensorResponse": {...}}
}

// fetchSpec fetches and parses the Swagger JSON spec from a service.
// Each service exposes its generated spec at /swagger/doc.json.
func fetchSpec(ctx context.Context, client *httpclient.Client, url string) (swaggerSpec, error) {
	return httpclient.DoRequest[swaggerSpec](client, ctx, url+"/swagger/doc.json", http.MethodGet, nil, nil)
}

// MergedSwaggerSpec returns a handler that fetches the Swagger spec from each
// service, merges them into a single spec, and returns it as JSON.
//
// Each service's paths are prefixed with its own basePath before merging,
// so "/sensors" from device-service becomes "/api/v1/device/sensors" in the
// merged output. This prevents path conflicts between services.
//
// The merged spec is served at GET /swagger/doc.json and consumed by the
// Swagger UI at GET /swagger/.
func MergedSwaggerSpec(deviceSvcURL, collSvcURL, authSvcURL string) http.HandlerFunc {
	// Client is created once and reused across requests
	client := httpclient.New()
	return func(w http.ResponseWriter, r *http.Request) {

		// Fetch each service's generated spec
		deviceSpec, err := fetchSpec(r.Context(), client, deviceSvcURL)
		if err != nil {
			http.Error(w, "failed to fetch device spec", http.StatusInternalServerError)
			return
		}

		collSpec, err := fetchSpec(r.Context(), client, collSvcURL)
		if err != nil {
			http.Error(w, "failed to fetch collection spec", http.StatusInternalServerError)
			return
		}

		authSpec, err := fetchSpec(r.Context(), client, authSvcURL)
		if err != nil {
			http.Error(w, "failed to fetch auth spec", http.StatusInternalServerError)
			return
		}

		// Build the merged spec with a unified info block.
		// Host is taken from the incoming request so it works in any environment.
		merged := swaggerSpec{
			Swagger:     "2.0",
			Info:        map[string]any{"title": "Innoveria IoT API", "version": "1.0"},
			Host:        r.Host,
			BasePath:    "/",
			Paths:       map[string]any{},
			Definitions: map[string]any{},
		}

		// Merge paths — prepend each service's basePath so routes are fully qualified
		for path, val := range deviceSpec.Paths {
			merged.Paths[deviceSpec.BasePath+path] = val
		}
		for path, val := range collSpec.Paths {
			merged.Paths[collSpec.BasePath+path] = val
		}
		for path, val := range authSpec.Paths {
			merged.Paths[authSpec.BasePath+path] = val
		}

		// Merge definitions (the request/response DTO schemas)
		maps.Copy(merged.Definitions, deviceSpec.Definitions)
		maps.Copy(merged.Definitions, collSpec.Definitions)
		maps.Copy(merged.Definitions, authSpec.Definitions)

		// Encode and return the merged spec.
		if err := json.Encode(w, http.StatusOK, merged); err != nil {
			http.Error(w, "failed to encode merged swagger spec", http.StatusInternalServerError)
		}
	}
}
