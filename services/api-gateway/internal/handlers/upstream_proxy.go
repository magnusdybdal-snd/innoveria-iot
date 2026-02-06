package handlers

import (
	"innoveria-iot/api-gateway/internal/proxy"
	"net/http"
	"net/url"
)

func NewUpstreamProxy(baseURL string) (http.Handler, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	_ = proxy.NewReverseProxy(base)

	return nil, nil
}
