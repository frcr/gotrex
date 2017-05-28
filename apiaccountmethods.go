package gotrex

import (
	"errors"
	"strconv"
)

/*
This file contains implementations of
Bittrex private api methods concerning account actions.
For objects' structure check the api/apiobjects.go file
*/

/*
/account/getbalances?apikey=API_KEY
/account/getbalance?apikey=API_KEY&currency=BTC
/account/getdepositaddress?apikey=API_KEY&currency=VTC
/account/withdraw?apikey=API_KEY&currency=EAC&quantity=20.40&address=EAC_ADDRESS
/account/getorder&uuid=0cb4c4e4-bdc7-4e13-8c13-430e587d2cc1
/account/getorderhistory
/account/getwithdrawalhistory?currency=BTC
/account/getdeposithistory?currency=BTC
*/

func (c *Connector) GetBalances() (*[]Balance, error) {
	var b []Balance
	err := c.UseMethod("/account/getbalances", &b)
	return &b, err
}

func (c *Connector) GetBalance(cur string) (*Balance, error) {
	if len(cur) == 0 {
		return nil, errors.New("Currency not specified")
	}
	var b Balance
	method := "/account/getbalance?currency=" + cur
	err := c.UseMethod(method, &b)
	return &b, err
}

func (c *Connector) GetDepositAddress(cur string) (*Address, error) {
	if len(cur) == 0 {
		return nil, errors.New("Currency not specified")
	}
	var a Address
	method := "/account/getdepositaddress?currency=" + cur
	err := c.UseMethod(method, &a)
	return &a, err
}

func (c *Connector) Withdraw(cur, addr, paymentId string, q float64) (*Uuid, error) {
	if len(cur) == 0 || len(addr) == 0 {
		return nil, errors.New("Currency and address are requred fields")
	}

	if q == 0.0 {
		return nil, errors.New("Can't withdraw 0.0 coins")
	}
	method := "/account/withdraw?currency=" + cur +
		"&quantity=" + strconv.FormatFloat(q, 'f', 8, 64) +
		"&address=" + addr
	if len(paymentId) != 0 {
		method += "&paymentid=" + paymentId
	}
	var u Uuid
	err := c.UseMethod(method, &u)
	return &u, err
}

func (c *Connector) GetOrder(uuid string) (*OrderGetter, error) {
	if len(uuid) == 0 {
		return nil, errors.New("No order uuid provided")
	}
	var og OrderGetter
	err := c.UseMethod("/account/getorder&uuid="+uuid, &og)
	return &og, err
}

func (c *Connector) GetOrderHistory(market string) (*[]OldOrder, error) {
	method := "/account/getorderhistory"
	if len(market) != 0 {
		method += "?market=" + market
	}
	var oo []OldOrder
	err := c.UseMethod(method, &oo)
	return &oo, err
}

func (c *Connector) GetWithdrawalHistory(cur string) (*[]Withdrawal, error) {
	method := "/account/getwithdrawalhistory"
	if len(cur) != 0 {
		method += "?currency=" + cur
	}
	var wdr []Withdrawal
	err := c.UseMethod(method, &wdr)
	return &wdr, err
}

func (c *Connector) GetDepositHistory(cur string) (*[]Deposit, error) {
	method := "/account/getdeposithistory"
	if len(cur) != 0 {
		method += "?currency=" + cur
	}
	var dps []Deposit
	err := c.UseMethod(method, &dps)
	return &dps, err
}
