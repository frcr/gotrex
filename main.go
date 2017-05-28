package gotrex

import (
	"errors"
	"net/http"
)

func NewPublicConnector() (*PublicConnector, error) {
	pc := PublicConnector{
		http.Client{},
		APIVER,
		ENDPOINT,
		PROTOCOL,
		TIMEOUT}
	return &pc, nil
}

func NewConnector(k, s string) (*Connector, error) {
	if len(k) == 0 || len(s) == 0 {
		return nil, errors.New("Error: either api key or secret is nil.")
	}
	p, _ := NewPublicConnector()
	c := Connector{*p, k, s}
	return &c, nil
}

func (c *PublicConnector) UseMethod(method string, v interface{}) error {
	req, _ := c.makeRequest(method)
	return c.decodePayload(req, v)
}

func (c *Connector) UseMethod(method string, v interface{}) error {
	req, _ := c.makeRequest(method)
	req = c.encrypt(req)
	return c.decodePayload(req, v)
}
