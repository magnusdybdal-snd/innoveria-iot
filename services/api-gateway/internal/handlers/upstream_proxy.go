package handlers

import (
	"net/http"
	"net/url"

	"innoveria-iot/api-gateway/internal/proxy"
)

func NewUpstreamProxy(baseURL string) (http.Handler, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	rp := proxy.NewReverseProxy(base)

	return rp, nil
}
