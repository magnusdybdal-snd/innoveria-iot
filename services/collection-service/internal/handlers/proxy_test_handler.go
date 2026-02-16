package handlers

import (
	"innoveria-iot/pkg/json"
	"net/http"
)

// TODO: Remove, this is used to test the proxy from api gateway
func HelloProxy(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"hello": "world",
	}
	_ = json.Encode(w, http.StatusOK, resp)
}
