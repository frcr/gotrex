package gotrex

type Market struct {
	MarketCurrency     string   `json:"MarketCurrency"`
	BaseCurrency       string   `json:"BaseCurrency"`
	MarketCurrencyLong string   `json:"MarketCurrencyLong"`
	BaseCurrencyLong   string   `json:"BaseCurrencyLong"`
	MinTradeSize       float64  `json:"MinTradeSize"`
	MarketName         string   `json:"MarketName"`
	IsActive           bool     `json:"IsActive"`
	Created            jsonTime `json:"Created"`
}

type Currency struct {
	Currency        string  `json:"Currency"`
	CurrencyLong    string  `json:"CurrencyLong"`
	MinConfirmation int     `json:"MinConfirmation"`
	TxFee           float64 `json:"TxFee"`
	IsActive        bool    `json:"IsActive"`
	CoinType        string  `json:"CoinType"`
	BaseAddress     string  `json:"BaseAddress"`
}

type Ticker struct {
	Bid  float64 `json:"Bid"`
	Ask  float64 `json:"Ask"`
	Last float64 `json:"Last"`
}

type MarketSummary struct {
	MarketName        string   `json:"MarketName"`
	High              float64  `json:"High"`
	Low               float64  `json:"Low"`
	Volume            float64  `json:"Volume"`
	Last              float64  `json:"Last"`
	BaseVolume        float64  `json:"BaseVolume"`
	TimeStamp         string   `json:"TimeStamp"`
	Bid               float64  `json:"Bid"`
	Ask               float64  `json:"Ask"`
	OpenBuyOrders     int      `json:"OpenBuyOrders"`
	OpenSellOrders    int      `json:"OpenSellOrders"`
	PrevDay           float64  `json:"PrevDay"`
	Created           jsonTime `json:"Created"`
	DisplayMarketName string   `json:"DisplayMarketName"`
}

type OrderBook struct {
	Buy  []OrderBookEntry `json:"buy"`
	Sell []OrderBookEntry `json:"sell"`
}

type OrderBookEntry struct {
	Quantity float64 `json:"Quantity"`
	Rate     float64 `json:"Rate"`
}

type Trade struct {
	Id        int      `json:"Id"`
	Timestamp jsonTime `json:"TimeStamp"`
	Quantity  float64  `json:"Quantity"`
	Price     float64  `json:"Price"`
	Total     float64  `json:"Total"`
	FillType  string   `json:"FillType"`
	OrderType string   `json:"OrderType"`
}

type Uuid struct {
	Id string `json:"Uuid"`
}

type Order struct {
	Uuid              string   `json:"Uuid"`
	OrderUuid         string   `json:"OrderUuid"`
	Exchange          string   `json:"Exchange"`
	OrderType         string   `json:"OrderType"`
	Quantity          float64  `json:"Quantity"`
	QuantityRemaining float64  `json:"QuantityRemaining"`
	Limit             float64  `json:"Limit"`
	CommissionPaid    float64  `json:"CommissionPaid"`
	Price             float64  `json:"Price"`
	PricePerUnit      string   `json:"PricePerUnit"`
	Opened            jsonTime `json:"Opened"`
	Closed            jsonTime `json:"Closed"`
	CancelInitiated   bool     `json:"CancelInitiated"`
	ImmediateOrCancel bool     `json:"ImmediateOrCancel"`
	IsConditional     bool     `json:"IsConditional"`
	Condition         string   `json:"Condition"`
	ConditionTarget   string   `json:"ConditionTarget"`
}

type Balance struct {
	Currency      string  `json:"Currency"`
	Balance       float64 `json:"Balance"`
	Available     float64 `json:"Available"`
	Pending       float64 `json:"Pending"`
	CryptoAddress string  `json:"CryptoAddress"`
	Requested     bool    `json:"Requested"`
	Uuid          string  `json:"Uuid"`
}

type Address struct {
	Currency string `json:"Currency"`
	Address  string `json:"Address"`
}

// The method "/account/getorder" requires special handling
type OrderGetter struct {
	AccountId                  string  `json:"AccountId"`
	OrderUuid                  string  `json:"OrderUuid"`
	Exchange                   string  `json:"Exchange"`
	Type                       string  `json:"Type"`
	Quantity                   float64 `json:"Quantity"`
	QuantityRemaining          float64 `json:"QuantityRemaining"`
	Limit                      float64 `json:"Limit"`
	Reserved                   float64 `json:"Reserved"`
	ReserveRemaining           float64 `json:"ReserveRemaining"`
	CommissionReserved         float64 `json:"CommissionReserved"`
	CommissionReserveRemaining float64 `json:"CommissionReserveRemaining"`
	CommissionPaid             float64 `json:"CommissionPaid"`
	Price                      float64 `json:"Price"`
	PricePerUnit               float64 `json:"PricePerUnit"`
	Opened                     string  `json:"Opened"`
	Closed                     string  `json:"Closed"`
	IsOpen                     bool    `json:"IsOpen"`
	Sentinel                   string  `json:"Sentinel"`
	CancelInitiated            bool    `json:"CancelInitiated"`
	ImmediateOrCancel          bool    `json:"ImmediateOrCancel"`
	IsConditional              bool    `json:"IsConditional"`
	Condition                  string  `json:"Condition"`
	ConditionTarget            string  `json:"ConditionTarget"`
}

type OldOrder struct {
	OrderUuid         string   `json:"OrderUuid"`
	Exchange          string   `json:"Exchange"`
	TimeStamp         jsonTime `json:"TimeStamp"`
	OrderType         string   `json:"OrderType"`
	Limit             float64  `json:"Limit"`
	Quantity          float64  `json:"Quantity"`
	QuantityRemaining float64  `json:"QuantityRemaining"`
	Commission        float64  `json:"Commission"`
	Price             float64  `json:"Price"`
	PricePerUnit      float64  `json:"PricePerUnit"`
	IsConditional     bool     `json:"IsConditional"`
	Condition         string   `json:"Condition"`
	ConditionTarget   string   `json:"ConditionTarget"`
	ImmediateOrCancel bool     `json:"ImmediateOrCancel"`
}

type Withdrawal struct {
	PaymentUuid    string   `json:"PaymentUuid"`
	Currency       string   `json:"Currency"`
	Amount         float64  `json:"Amount"`
	Address        string   `json:"Address"`
	Opened         jsonTime `json:"Opened"`
	Authorized     bool     `json:"Authorized"`
	PendingPayment bool     `json:"PendingPayment"`
	TxCost         float64  `json:"TxCost"`
	TxId           string   `json:"TxId"`
	Canceled       bool     `json:"Canceled"`
	InvalidAddress bool     `json:"InvalidAddress"`
}

type Deposit struct {
	Id            int64    `json:"Id"`
	Amount        float64  `json:"Amount"`
	Currency      string   `json:"Currency"`
	Confirmations int      `json:"Confirmations"`
	LastUpdated   jsonTime `json:"LastUpdated"`
	TxId          string   `json:"TxId"`
	CryptoAddress string   `json:"CryptoAddress"`
}
