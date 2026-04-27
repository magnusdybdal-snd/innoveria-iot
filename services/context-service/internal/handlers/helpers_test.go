package handlers_test

import "net/http"

func withAuthHeaders(r *http.Request) *http.Request {
	r.Header.Set("X-Auth-User-Id", "u0000000-0000-0000-0000-000000000001")
	r.Header.Set("X-Auth-Company-Id", "a0000000-0000-0000-0000-000000000001")
	r.Header.Set("X-Auth-Role", "FACTORY_WORKER")
	return r
}
