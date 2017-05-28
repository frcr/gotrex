# GoTrex

Implementation of a wrapper for Bittrex API v1.1, written in Go.

### Installation

To install the package use the usual go commands, given you have your `$GOPATH` variable set
~~~
go get github.com/frcr/gotrex
~~~
Then just import it in your program and you're good to go:
~~~go
import(
	"github.com/frcr/gotrex"
)
~~~
### Usage

You can find the discription of API itself on the official [Bittrex site](https://bittrex.com/Home/Api).

The usage of the package revolves around two main `struct`s: `Connector` and `PublicConnector`

The `PublicConnector` is able to only handle the public methods, as it contains neither methods, nor API key to access the data that requires authorisation.

The `Connector` struct can handle any methods of the API.

To start using the API you need to get a new connector (for the `Connector` you will have to provide api key and secret):
~~~go
	publicOne, _ := gotrex.NewPublicConnector()
	secureOne, _ := gotrex.NewConnector(apikey, apisecret string)
~~~
You can check out the implementation of these functions in `main.go`. The implementation of the connectors is in the `connector.go`

Now you have two options:
1. You can call the methods described in
	*	`apipublicmethods.go`
	*	`apimarketmethods.go`
	*	`apiaccountmethods.go`
2. You can use the general `UseMethod()` method

The first option allows you to use go methods like
~~~go
sumry, _ := publicOne.GetMarketSummary("BTC-DOGE")
~~~
but with the second you have more freedom forming your own urls, though it *is* functionally equivalent:
~~~go
var sumry MarketSummary 
err := publicOne.UseMethod("/public/getmarketsummary?market=BTC-DOGE", &sumry)
~~~
Of course, you'll also have to check the correctness of your formed paths and arguments, as `UseMethod` assumes the provided method and arguments are correct. But it handles authorisation, so you don't have to add `"&apikey="`, `"&nonce="` or do any encryption and signing if you use it with `Connector`:
~~~go
c, err := gotrex.NewConnector(key, secret)
method := "/market/cancel?uuid=" + uuid
err = c.UseMethod(method)
~~~
