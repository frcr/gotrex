package gotrex

import (
	"errors"
)

/*
This file contains implementations of
Bittrex private api methods concerning market actions.
For objects' structure check the api/apiobjects.go file
*/

/*

/market/buylimit?apikey=API_KEY&market=BTC-LTC&quantity=1.2&rate=1.3
/market/selllimit?apikey=API_KEY&market=BTC-LTC&quantity=1.2&rate=1.3
/market/cancel?apikey=API_KEY&uuid=ORDER_UUID
/market/getopenorders?apikey=API_KEY&market[optional]=BTC-LTC

*/

func (c *Connector) BuyLimit(market string, q, r float64) (*Uuid, error) {
	return c.limitOrder("buy", market, q, r)
}

func (c *Connector) SellLimit(market string, q, r float64) (*Uuid, error) {
	return c.limitOrder("sell", market, q, r)
}

func (c *Connector) Cancel(uuid string) error {
	if len(uuid) == 0 {
		return errors.New("Empty UUID")
	}
	method := "/market/cancel?uuid=" + uuid
	var i []interface{}
	err := c.UseMethod(method, &i)
	return err
}

func (c *Connector) GetOpenOrders(market string) (*[]Order, error) {
	method := "/market/getopenorders"
	if len(market) != 0 {
		method += "?market=" + market
	}
	var rez []Order
	err := c.UseMethod(method, &rez)
	return &rez, err
}
