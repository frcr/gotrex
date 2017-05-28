package gotrex

import (
	"errors"
	"strconv"
)

/*
This file contains implementations of Bittrex public api methods.
For objects' structure check the api/apiobjects.go file
*/

/*
//public

/public/getmarkets
/public/getcurrencies
/public/getticker?market=BTC-LTC
/public/getmarketsummaries
/public/getmarketsummary?market=BTC-LTC
/public/getorderbook?market=BTC-LTC&type=both&depth=50
		type=[buy, sell, both]
		depth <= 50
/public/getmarkethistory?market=BTC-DOGE

*/

func (c *PublicConnector) GetMarkets() (*[]Market, error) {
	var marketList []Market
	err := c.UseMethod("/public/getmarkets", &marketList)
	return &marketList, err
}

func (c *PublicConnector) GetCurrencies() (*[]Currency, error) {
	var curList []Currency
	err := c.UseMethod("/public/getcurrencies", &curList)
	return &curList, err
}

func (c *PublicConnector) GetTicker(t string) (*Ticker, error) {
	/*
		t: required field, example:"BTC-LTC"
		For full list of currency pairs, use GetMarketSummaries()
	*/
	if len(t) == 0 {
		return nil, errors.New("No market tag to get ticker from")
	}
	var rezult Ticker
	method := "/public/getticker?market=" + t
	err := c.UseMethod(method, &rezult)
	return &rezult, err
}

func (c *PublicConnector) GetMarketSummaries() (*[]MarketSummary, error) {
	var ms []MarketSummary
	err := c.UseMethod("/public/getmarketsummaries", &ms)
	return &ms, err
}

func (c *PublicConnector) GetMarketSummary(t string) (*MarketSummary, error) {
	/*
		t: required field, example:"BTC-LTC"
		For full list of currency pairs, use GetMarketSummaries()
	*/
	if len(t) == 0 {
		return nil, errors.New("No market tag to get summary of")
	}
	var rezult MarketSummary
	method := "/public/getmarketsummary?market=" + t
	err := c.UseMethod(method, &rezult)
	return &rezult, err
}

func (c *PublicConnector) GetOrderBook(market, orderTypes string, depth int) (*OrderBook, error) {
	/*
		market: required field, example:"BTC-LTC"
		For full list of currency pairs, use GetMarketSummaries()

		orderTypes: required field, possible values: "buy", "sell", "both"

		depth: if not in range [1, 2, 3,... 50], defaults to 20 on the server side.

	*/
	if len(market) == 0 || len(orderTypes) == 0 {
		return nil, errors.New("Market tag and type of order are required fields")
	}
	method := "/public/getorderbook?market=" + market +
		"&type=" + orderTypes
	if depth > 1 && depth <= 50 {
		method += "&depth=" + strconv.Itoa(depth)
	}
	var rezult OrderBook
	err := c.UseMethod(method, &rezult)
	return &rezult, err
}

func (c *PublicConnector) GetMarketHistory(tag string) (*[]Trade, error) {
	if len(tag) == 0 {
		return nil, errors.New("No market to get history of")
	}
	var tr []Trade
	err := c.UseMethod("/public/getmarkethistory?market="+tag, &tr)
	return &tr, err
}
