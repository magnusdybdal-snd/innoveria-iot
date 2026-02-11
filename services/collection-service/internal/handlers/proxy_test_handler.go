package handlers

import (
	"net/http"
	"innoveria-iot/pkg/json"
)

func HelloProxy(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string {
		"hello": "world",
	}
	_ = json.Encode(w, http.StatusOK, resp)
}
