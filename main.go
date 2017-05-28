package gotrex

import (
	"net/http"
)

func NewPublicConnector() (*PublicConnector, error) {
	return *PublicConnector{http.Client{},
		ApiVer:   APIVER,
		Endpoint: ENDPOINT,
		Protocol: PROTOCOL,
		Timeout:  TIMEOUT}, nil
}

func NewConnector(k, s string) (*Connector, error) {
	if len(k) == 0 || len(s) == 0 {
		return nil, errors.New("Error: either api key or secret is nil.")
	}
	return *Connector{*NewPublicConnector(), ApiKey: k, ApiSec: s}, nil
}

func (c *PublicConnector) UseMethod(method string, v *interface{}) error {
	req := c.makeRequest(method)
	return c.decodePayload(req, v)
}

func (c *Connector) UseMethod(method string, v *interface{}) error {
	req := c.makeRequest(method)
	req = c.encrypt(req)
	return c.decodePayload(req, v)
}
