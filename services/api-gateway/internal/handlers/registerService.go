package handlers

import (
	"net/http"

	"innoveria-iot/pkg/json"
)

// RegisterServiceError TODO(@vinjar): add proper documentation.
func RegisterServiceError(mux *http.ServeMux, route, name string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_ = json.Encode(w, http.StatusServiceUnavailable, map[string]string{
			"error": name + " unavailable",
		})
	}

	mux.HandleFunc(route, handler)
	mux.HandleFunc(route+"/", handler)
}

// RegisterServiceInfo TODO(@vinjar): add proper documentation.
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

// RegisterProxyService TODO(@vinjar): add proper documentation.
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
	for _, path := range paths {
		full := route + path
		mux.Handle(full, proxy)
		mux.Handle(full+"/", proxy)
	}
}
