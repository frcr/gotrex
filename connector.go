package gotrex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type PublicConnector struct {
	http.Client
	ApiVer   string
	Endpoint string
	Protocol string
	Timeout  int
}
type Connector struct {
	PublicConnector
	ApiKey string
	ApiSec string
}

type jsonChecker struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

func (c *PublicConnector) makeRequest(method string) (*http.Request, error) {
	var generatedUrl = "" +
		c.Protocol + "://" +
		c.Endpoint + "/" +
		c.ApiVer +
		method
	req, err := http.NewRequest("GET", generatedUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	return req, nil
}

func (c *Connector) encrypt(req *http.Request) *http.Request {
	nonce := time.Now().UnixNano()
	q := req.URL.Query()
	q.Set("apikey", c.ApiKey)
	q.Set("nonce", fmt.Sprintf("%d", nonce))
	req.URL.RawQuery = q.Encode()
	algo := hmac.New(sha512.New, []byte(c.ApiSec))
	algo.Write([]byte(req.URL.String()))
	signature := hex.EncodeToString(algo.Sum(nil))
	req.Header.Add("apisign", signature)
	return req
}

func (c *PublicConnector) launch(req *http.Request) (*http.Response, error) {
	timeout := time.NewTimer(TIMEOUT * time.Second)

	type pack struct {
		r *http.Response
		e error
	}
	done := make(chan pack, 1)
	go func() {
		response, err := c.Do(req)
		tmp := pack{response, err}
		done <- tmp
	}()
	// Wait for the read or the timeout
	select {
	case ret := <-done:
		return ret.r, ret.e
	case <-timeout.C:
		return nil, errors.New("timeout on reading data from Bittrex API")
	}
}

func (c *PublicConnector) decodePayload(req *http.Request, v interface{}) error {
	response, err := c.launch(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	bytestring, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var jc jsonChecker
	if err = json.Unmarshal(bytestring, &jc); err != nil {
		return err
	}
	if err = jc.check(); err != nil {
		return err
	}
	err = json.Unmarshal(jc.Result, v)
	return err
}

func (jc *jsonChecker) check() error {
	if !jc.Success {
		return errors.New(jc.Message)
	}
	return nil
}

func (c *Connector) limitOrder(otype string, market string, q, r float64) (*Uuid, error) {
	if q == 0.0 || r == 0.0 {
		return nil, errors.New("Quantity and rate should be non-zero")
	}
	if len(market) == 0 {
		return nil, errors.New("Market tag is required")
	}
	var rezult Uuid
	method := "/market/" + otype + "limit?market=" +
		market +
		"&quantity=" +
		strconv.FormatFloat(q, 'f', 8, 64) +
		"&rate=" +
		strconv.FormatFloat(r, 'f', 8, 64)
	err := c.UseMethod(method, &rezult)
	return &rezult, err
}
