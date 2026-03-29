package monitor

import (
	"fmt"
	"innoveria-iot/pkg/httpclient"
)

type Client struct {
	host, lang, company string
	username, password  string
	httpClient          *httpclient.Client
	sessionID           string
}

// base returns the base url for accessing monitor erp
func (c *Client) base() string {
	return fmt.Sprintf("https://%s:8001/%s/%s", c.host, c.lang, c.company)
}

// loginUrl returns the url for authenticating with monitor erp
// it will return a session id which is used throughout the api calls
func (c *Client) loginUrl() string {
	return c.base() + "/login"
}

// apiUrl returns the endpoint for
func (c *Client) apiUrl(path string) string {
	return c.base() + "/api/v1/" + path
}

func (c *Client) ensureSession() error {
	return nil
}
