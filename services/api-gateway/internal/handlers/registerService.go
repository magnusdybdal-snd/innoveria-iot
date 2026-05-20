package handlers

import (
	"net/http"

	"innoveria-iot/pkg/json"
)

// NotFound returns a 404 response. Use this to explicitly block specific routes.
func NotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// RegisterServiceError registers a handler on route that always returns 503, used when an upstream service could not be reached at startup.
func RegisterServiceError(mux *http.ServeMux, route, name string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_ = json.Encode(w, http.StatusServiceUnavailable, map[string]string{
			"error": name + " unavailable",
		})
	}

	mux.HandleFunc(route, handler)
	mux.HandleFunc(route+"/", handler)
}

// RegisterServiceInfo registers a handler on route that returns a JSON summary of the service name and its available paths.
func RegisterServiceInfo(mux *http.ServeMux, route, name string, paths []string) {
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		full := make([]string, len(paths))
		for i, p := range paths {
			full[i] = route + p
		}
		_ = json.Encode(w, http.StatusOK, map[string]any{
			"service": name,
			"routes":  full,
		})
	})
}

// RegisterProxyService registers reverse-proxy handlers for each path of an upstream service.
func RegisterProxyService(
	mux *http.ServeMux,
	route string,
	serviceName string,
	baseUrl string,
	paths []string,
) {
	proxy, err := NewUpstreamProxy(baseUrl)
	if err != nil {
		RegisterServiceError(mux, route, serviceName)
		return
	}
	RegisterServiceInfo(mux, route, serviceName, paths)
	seen := make(map[string]struct{}, len(paths))
	for _, path := range paths {
		if _, ok := seen[path]; ok {
			continue
		}
		seen[path] = struct{}{}
		full := route + path
		mux.Handle(full, proxy)
		mux.Handle(full+"/", proxy)
	}
}
