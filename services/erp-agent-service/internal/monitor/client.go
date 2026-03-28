package monitor

import "innoveria-iot/pkg/httpclient"

type Client struct {
	host, lang, company string
	username, password  string
	httpClient          *httpclient.Client
	sessionID           string
}

func (c *Client) ensureSession() error
