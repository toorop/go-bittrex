// Package Bittrex is an implementation of the Biitrex API in Golang.
package bittrex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/shopspring/decimal"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	API_BASE    = "https://api.bittrex.com/" // Bittrex API endpoint
	API_VERSION = "v3"
	WS_BASE     = "socket.bittrex.com" // Bittrex WS API endpoint
	WS_HUB      = "CoreHub"            // SignalR main hub
)

// New returns an instantiated bittrex struct
func New(apiKey, apiSecret string) *Bittrex {
	client := NewClient(apiKey, apiSecret)
	return &Bittrex{client}
}

// NewWithCustomHttpClient returns an instantiated bittrex struct with custom http client
func NewWithCustomHttpClient(apiKey, apiSecret string, httpClient *http.Client) *Bittrex {
	client := NewClientWithCustomHttpConfig(apiKey, apiSecret, httpClient)
	return &Bittrex{client}
}

// NewWithCustomTimeout returns an instantiated bittrex struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *Bittrex {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &Bittrex{client}
}

// handleErr gets JSON response from Bittrex API en deal with error
func handleErr(r jsonResponse) error {
	if !r.Success {
		return errors.New(r.Message)
	}
	return nil
}

// bittrex represent a bittrex client
type Bittrex struct {
	client *client
}

// set enable/disable http request/response dump
func (c *Bittrex) SetDebug(enable bool) {
	c.client.debug = enable
}

// GetDistribution is used to get the distribution.
func (b *Bittrex) GetDistribution(market string) (distribution Distribution, err error) {
	r, err := b.client.do("GET", "https://bittrex.com/Api/v2.0/pub/currency/GetBalanceDistribution?currencyName="+strings.ToUpper(market), "", false)
	if err != nil {
		return
	}

	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}

	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &distribution)
	return

}

// GetCurrencies is used to get all supported currencies at Bittrex along with other meta data.
func (b *Bittrex) GetCurrencies() (currencies []CurrencyV3, err error) {
	r, err := b.client.do("GET", "currencies", "", false)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &currencies)
	return
}

// GetCurrency is used to get information of a single currency at Bittrex along with other meta data.
func (b *Bittrex) GetCurrency(symbol string) (currencies CurrencyV3, err error) {
	r, err := b.client.do("GET", "currencies/"+symbol, "", false)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &currencies)
	return
}

// GetMarkets is used to get the open and available trading markets at Bittrex along with other meta data.
func (b *Bittrex) GetMarkets() (markets []MarketV3, err error) {
	r, err := b.client.do("GET", "markets", "", false)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &markets)
	return
}

// GetTicker is used to get the current ticker values for a market, if none is specified, returns info for all.
func (b *Bittrex) GetTicker(market string) (ticker []TickerV3, err error) {
	market = strings.ToUpper(market)
	var endpoint string
	if market == "" {
		endpoint = "markets/tickers"
	} else {
		endpoint = "markets/" + market + "/ticker"
	}

	r, err := b.client.do("GET", endpoint, "", false)
	if err != nil {
		return
	}
	if market == "" {
		err = json.Unmarshal(r, &ticker)
	} else {
		var t TickerV3
		err = json.Unmarshal(r, &t)
		if err != nil {
			return
		}
		ticker = append(ticker, t)
	}

	return
}

// GetMarketSummaries is used to get the last 24 hour summary of all active exchanges
func (b *Bittrex) GetMarketSummaries() (marketSummaries []MarketSummaryV3, err error) {
	r, err := b.client.do("GET", "markets/summaries", "", false)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &marketSummaries)
	return
}

// GetMarketSummary is used to get the last 24 hour summary for a given market
func (b *Bittrex) GetMarketSummary(market string) (marketSummary MarketSummaryV3, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("markets/%s/summary", strings.ToUpper(market)), "", false)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &marketSummary)
	return
}

// GetOrderBook is used to get retrieve the orderbook for a given market
// market: a string literal for the market (ex: BTC-LTC)
// cat: buy, sell or both to identify the type of orderbook to return.
func (b *Bittrex) GetOrderBook(market string, depth int32, cat string) (orderBook OrderBookV3, err error) {
	if cat != "buy" && cat != "sell" && cat != "both" {
		cat = "both"
	}

	r, err := b.client.do("GET", fmt.Sprintf("markets/%s/orderbook?depth=%s", strings.ToUpper(market), strconv.Itoa(int(depth))), "", false)
	if err != nil {
		return
	}

	var auxOrderBook OrderBookV3
	// TODO Verify Ask and Bid logic is OK
	err = json.Unmarshal(r, &orderBook)

	if cat == "both" {
		return
	}

	if cat == "buy" {
		auxOrderBook.Ask = orderBook.Ask
	} else if cat == "sell" {
		auxOrderBook.Bid = orderBook.Bid
	}
	orderBook = auxOrderBook
	return
}

// GetOrderBookBuySell is used to get retrieve the buy or sell side of an orderbook for a given market
// market: a string literal for the market (ex: BTC-LTC)
// cat: buy or sell to identify the type of orderbook to return.
func (b *Bittrex) GetOrderBookBuySell(market string, depth int32, cat string) (orderb []OrderbV3, err error) {
	if cat != "buy" && cat != "sell" {
		cat = "buy"
	}

	r, err := b.client.do("GET", fmt.Sprintf("markets/%s/orderbook?depth=%s", strings.ToUpper(market), strconv.Itoa(int(depth))), "", false)
	if err != nil {
		return
	}
	var orderBook OrderBookV3
	err = json.Unmarshal(r, &orderBook)
	if err != nil {
		return
	}
	if cat == "buy" {
		orderb = orderBook.Ask
	} else if cat == "sell" {
		orderb = orderBook.Bid
	}
	return
}

// GetMarketHistory is used to retrieve the latest trades that have occured for a specific market.
// market a string literal for the market (ex: BTC-LTC)
func (b *Bittrex) GetMarketHistory(market string) (trades []TradeV3, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("markets/%s/trades", strings.ToUpper(market)), "", false)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &trades)
	return
}

// Market

// BuyLimit is used to place a limited buy order in a specific market.
func (b *Bittrex) BuyLimit(market string, quantity, rate decimal.Decimal) (uuid string, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("market/buylimit?market=%s&quantity=%s&rate=%s", market, quantity, rate), "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var u Uuid
	err = json.Unmarshal(response.Result, &u)
	uuid = u.Id
	return
}

// CreateOrder is used to create any type of supported order.
func (b *Bittrex) CreateOrder(params CreateOrderParams) (order OrderV3, err error) {
	// TODO Preprocessor
	if params.Type == "" || params.MarketSymbol == "" || params.Direction == "" || params.TimeInForce == "" {
		// Check for missing parameters
		return OrderV3{}, ERR_ORDER_MISSING_PARAMETERS
	}
	var finalParams = CreateOrderParams{
		MarketSymbol:  params.MarketSymbol,
		Direction:     params.Direction,
		Type:          params.Type,
		Quantity:      decimal.Decimal{},
		Ceiling:       decimal.Decimal{},
		Limit:         decimal.Decimal{},
		TimeInForce:   params.TimeInForce,
		ClientOrderID: params.ClientOrderID,
		UseAwards:     params.UseAwards,
	}

	switch params.Type {
	case MARKET:
		finalParams.Quantity = params.Quantity
	case LIMIT:
		finalParams.Limit = params.Limit
		finalParams.Quantity = params.Quantity
	case CEILING_LIMIT:
		finalParams.Ceiling = params.Ceiling
	case CEILING_MARKET:
		finalParams.Ceiling = params.Ceiling

	}

	payload, err := json.Marshal(finalParams)
	if err != nil {
		return
	}
	r, err := b.client.do("POST", fmt.Sprintf("orders"), string(payload), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &order)
	return
}

// CancelOrder is used to cancel a buy or sell order.
func (b *Bittrex) CancelOrder(orderID string) (order OrderV3, err error) {
	r, err := b.client.do("DELETE", "orders/"+orderID, "", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &order)
	return
}

// GetClosedOrders returns orders that you currently have opened.
// If market is set to "all", GetClosedOrders return all orders
// If market is set to a specific order, GetClosedOrders return orders for this market
func (b *Bittrex) GetClosedOrders(market string) (closedOrders []OrderV3, err error) {
	resource := "orders/closed"
	if market == "" {
		market = "all"
	}
	if market != "all" {
		resource += "?marketSymbol=" + strings.ToUpper(market)
	}
	r, err := b.client.do("GET", resource, "", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &closedOrders)
	return
}

// GetOpenOrders returns orders that you currently have opened.
// If market is set to "all", GetOpenOrders return all orders
// If market is set to a specific order, GetOpenOrders return orders for this market
func (b *Bittrex) GetOpenOrders(market string) (openOrders []OrderV3, err error) {
	resource := "orders/open"
	if market == "" {
		market = "all"
	}
	if market != "all" {
		resource += "?marketSymbol=" + strings.ToUpper(market)
	}
	r, err := b.client.do("GET", resource, "", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &openOrders)
	return
}

// Account

// GetBalances is used to retrieve all balances from your account
func (b *Bittrex) GetBalances() (balances []BalanceV3, err error) {
	r, err := b.client.do("GET", "balances", "", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &balances)
	return
}

// Getbalance is used to retrieve the balance from your account for a specific currency.
// currency: a string literal for the currency (ex: LTC)
func (b *Bittrex) GetBalance(currency string) (balance Balance, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("balances/%s", strings.ToUpper(currency)), "", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &balance)
	return
}

// GetDepositAddress is sed to generate or retrieve an address for a specific currency.
// currency a string literal for the currency (ie. BTC)
func (b *Bittrex) GetDepositAddress(currency string) (address AddressV3, err error) {
	var addressParams = AddressParams{CurrencySymbol: currency}
	payload, err := json.Marshal(addressParams)
	if err != nil {
		return
	}
	r, err := b.client.do("GET", fmt.Sprintf("addresses/%s", currency), "", true)
	/* r, err := b.client.do("POST", "addresses", string(payload), true)
	if err != nil {
		return
	} */
	err = json.Unmarshal(r, &address)
	if err != nil {
		return
	}

	if address.CryptoAddress == "" {
		log.Println("needs to create new address")
		_, _ = b.client.do("POST", "addresses", string(payload), true)
		r, err = b.client.do("GET", fmt.Sprintf("addresses/%s", currency), "", true)
		if err != nil {
			return
		}
		err = json.Unmarshal(r, &address)
	}
	return
}

// Withdraw is used to withdraw funds from your account.
// address string the address where to send the funds.
// currency string literal for the currency (ie. BTC)
// quantity decimal.Decimal the quantity of coins to withdraw
// tag string an optional name for the withdrawal address
func (b *Bittrex) Withdraw(address, currency string, quantity decimal.Decimal, tag string) (withdraw WithdrawalV3, err error) {
	if address == "" || currency == "" || quantity.LessThan(decimal.NewFromFloat(0.0)) {
		return withdraw, ERR_WITHDRAWAL_MISSING_PARAMETERS
	}
	var params = WithdrawalParams{
		CurrencySymbol:   currency,
		Quantity:         quantity.String(),
		CryptoAddress:    address,
		CryptoAddressTag: "",
	}
	payload, err := json.Marshal(params)
	r, err := b.client.do("POST", "withdrawals", string(payload), true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &withdraw)
	return
}

// GetOpenWithdrawals is used to retrieve your open withdrawal history
// currency string a string literal for the currency (ie. BTC). If set to "", will return for all currencies
func (b *Bittrex) GetOpenWithdrawals(currency string, status WithdrawalStatus) (withdrawals []WithdrawalV3, err error) {
	var params = WithdrawalHistoryParams{
		Status:         string(status),
		CurrencySymbol: strings.ToUpper(currency),
	}
	v, _ := query.Values(params)
	queryParams := v.Encode()
	resource := "withdrawals/open"
	if len(queryParams) != 0 {
		resource += "?"
	}
	r, err := b.client.do("GET", resource+queryParams, "", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &withdrawals)
	return
}

// GetClosedWithdrawals is used to retrieve your closed withdrawal history
// currency string a string literal for the currency (ie. BTC). If set to "all", will return for all currencies
// TODO Add more parameters according to https://bittrex.github.io/api/v3#operation--withdrawals-closed-get
func (b *Bittrex) GetClosedWithdrawals(currency string, status WithdrawalStatus) (withdrawals []WithdrawalV3, err error) {
	var params = WithdrawalHistoryParams{
		Status:         string(status),
		CurrencySymbol: strings.ToUpper(currency),
	}
	v, _ := query.Values(params)
	queryParams := v.Encode()
	resource := "withdrawals/closed"
	if len(queryParams) != 0 {
		resource += "?"
	}
	r, err := b.client.do("GET", resource+queryParams, "", true)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &withdrawals)
	return
}

// GetWithdrawalByTxId is used to retrieve information of a withdrawal by txid.
// txid string a string literal for the on chain transaction id.
func (b *Bittrex) GetWithdrawalByTxId(txid string) (withdrawal WithdrawalV3, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("withdrawals/ByTxId/%s", txid), "", true)
	if err != nil {
		return
	}
	var withdrawals []WithdrawalV3
	err = json.Unmarshal(r, &withdrawals)
	if err != nil {
		return
	}
	if len(withdrawals) > 0 {
		return withdrawals[0], nil
	}
	return
}

// GetDepositHistory is used to retrieve your deposit history
// currency string a string literal for the currency (ie. BTC). If set to "all", will return for all currencies
func (b *Bittrex) GetDepositHistory(txid string) (deposits []DepositV3, err error) {
	resource := fmt.Sprintf("deposits/ByTxId/%s", txid)
	r, err := b.client.do("GET", resource, "", true)
	if err != nil {
		return
	}
	// var response json.RawMessage
	if err = json.Unmarshal(r, &deposits); err != nil {
		return
	}
	/*if err = handleErr(response); err != nil {
		return
	}*/
	// err = json.Unmarshal(response, &deposits)
	return
}

func (b *Bittrex) GetOrder(order_uuid string) (order Order2, err error) {

	resource := "account/getorder?uuid=" + order_uuid

	r, err := b.client.do("GET", resource, "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &order)
	return
}

// GetTicks is used to get ticks history values for a market.
// Interval can be -> ["oneMin", "fiveMin", "thirtyMin", "hour", "day"]
func (b *Bittrex) GetTicks(market string, interval string) ([]Candle, error) {
	_, ok := CANDLE_INTERVALS[interval]
	if !ok {
		return nil, errors.New("wrong interval")
	}

	endpoint := fmt.Sprintf(
		"https://bittrex.com/Api/v2.0/pub/market/GetTicks?tickInterval=%s&marketName=%s&_=%d",
		interval, strings.ToUpper(market), rand.Int(),
	)
	r, err := b.client.do("GET", endpoint, "", false)
	if err != nil {
		return nil, fmt.Errorf("could not get market ticks: %v", err)
	}

	var response jsonResponse
	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	if err := handleErr(response); err != nil {
		return nil, err
	}
	var candles []Candle

	if err := json.Unmarshal(response.Result, &candles); err != nil {
		return nil, fmt.Errorf("could not unmarshal candles: %v", err)
	}

	return candles, nil
}

// GetLatestTick returns array with a single element latest candle object
func (b *Bittrex) GetLatestTick(market string, interval string) ([]Candle, error) {
	_, ok := CANDLE_INTERVALS[interval]
	if !ok {
		return nil, errors.New("wrong interval")
	}

	endpoint := fmt.Sprintf(
		"https://bittrex.com/Api/v2.0/pub/market/GetLatestTick?tickInterval=%s&marketName=%s&_=%d",
		interval, strings.ToUpper(market), rand.Int(),
	)
	r, err := b.client.do("GET", endpoint, "", false)
	if err != nil {
		return nil, fmt.Errorf("could not get market ticks: %v", err)
	}

	var response jsonResponse
	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	if err := handleErr(response); err != nil {
		return nil, err
	}
	var candles []Candle

	if err := json.Unmarshal(response.Result, &candles); err != nil {
		return nil, fmt.Errorf("could not unmarshal candles: %v", err)
	}

	return candles, nil
}
