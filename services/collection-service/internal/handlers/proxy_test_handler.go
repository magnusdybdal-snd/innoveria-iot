package handlers

import (
	"innoveria-iot/pkg/json"
	"net/http"
)

func HelloProxy(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"hello": "world",
	}
	_ = json.Encode(w, http.StatusOK, resp)
}
