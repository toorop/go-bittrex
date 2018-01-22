// Package Bittrex is an implementation of the Biitrex API in Golang.
package bittrex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	API_BASE    = "https://bittrex.com/api/" // Bittrex API endpoint
	API_VERSION = "v1.1"
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
func (b *Bittrex) GetCurrencies() (currencies []Currency, err error) {
	r, err := b.client.do("GET", "public/getcurrencies", "", false)
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
	err = json.Unmarshal(response.Result, &currencies)
	return
}

// GetMarkets is used to get the open and available trading markets at Bittrex along with other meta data.
func (b *Bittrex) GetMarkets() (markets []Market, err error) {
	r, err := b.client.do("GET", "public/getmarkets", "", false)
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
	err = json.Unmarshal(response.Result, &markets)
	return
}

// GetTicker is used to get the current ticker values for a market.
func (b *Bittrex) GetTicker(market string) (ticker Ticker, err error) {
	r, err := b.client.do("GET", "public/getticker?market="+strings.ToUpper(market), "", false)
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
	err = json.Unmarshal(response.Result, &ticker)
	return
}

// GetMarketSummaries is used to get the last 24 hour summary of all active exchanges
func (b *Bittrex) GetMarketSummaries() (marketSummaries []MarketSummary, err error) {
	r, err := b.client.do("GET", "public/getmarketsummaries", "", false)
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
	err = json.Unmarshal(response.Result, &marketSummaries)
	return
}

// GetMarketSummary is used to get the last 24 hour summary for a given market
func (b *Bittrex) GetMarketSummary(market string) (marketSummary []MarketSummary, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("public/getmarketsummary?market=%s", strings.ToUpper(market)), "", false)
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
	err = json.Unmarshal(response.Result, &marketSummary)
	return
}

// GetOrderBook is used to get retrieve the orderbook for a given market
// market: a string literal for the market (ex: BTC-LTC)
// cat: buy, sell or both to identify the type of orderbook to return.
func (b *Bittrex) GetOrderBook(market, cat string) (orderBook OrderBook, err error) {
	if cat != "buy" && cat != "sell" && cat != "both" {
		cat = "both"
	}
	r, err := b.client.do("GET", fmt.Sprintf("public/getorderbook?market=%s&type=%s", strings.ToUpper(market), cat), "", false)
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

	if cat == "buy" {
		err = json.Unmarshal(response.Result, &orderBook.Buy)
	} else if cat == "sell" {
		err = json.Unmarshal(response.Result, &orderBook.Sell)
	} else {
		err = json.Unmarshal(response.Result, &orderBook)
	}

	return
}

// GetOrderBookBuySell is used to get retrieve the buy or sell side of an orderbook for a given market
// market: a string literal for the market (ex: BTC-LTC)
// cat: buy or sell to identify the type of orderbook to return.
func (b *Bittrex) GetOrderBookBuySell(market, cat string) (orderb []Orderb, err error) {
	if cat != "buy" && cat != "sell" {
		cat = "buy"
	}

	r, err := b.client.do("GET", fmt.Sprintf("public/getorderbook?market=%s&type=%s", strings.ToUpper(market), cat), "", false)
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
	err = json.Unmarshal(response.Result, &orderb)
	return
}

// GetMarketHistory is used to retrieve the latest trades that have occured for a specific market.
// market a string literal for the market (ex: BTC-LTC)
func (b *Bittrex) GetMarketHistory(market string) (trades []Trade, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("public/getmarkethistory?market=%s", strings.ToUpper(market)), "", false)
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
	err = json.Unmarshal(response.Result, &trades)
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

// SellLimit is used to place a limited sell order in a specific market.
func (b *Bittrex) SellLimit(market string, quantity, rate decimal.Decimal) (uuid string, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("market/selllimit?market=%s&quantity=%s&rate=%s", market, quantity, rate), "", true)
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

// CancelOrder is used to cancel a buy or sell order.
func (b *Bittrex) CancelOrder(orderID string) (err error) {
	r, err := b.client.do("GET", "market/cancel?uuid="+orderID, "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	err = handleErr(response)
	return
}

// GetOpenOrders returns orders that you currently have opened.
// If market is set to "all", GetOpenOrders return all orders
// If market is set to a specific order, GetOpenOrders return orders for this market
func (b *Bittrex) GetOpenOrders(market string) (openOrders []Order, err error) {
	resource := "market/getopenorders"
	if market != "all" {
		resource += "?market=" + strings.ToUpper(market)
	}
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
	err = json.Unmarshal(response.Result, &openOrders)
	return
}

// Account

// GetBalances is used to retrieve all balances from your account
func (b *Bittrex) GetBalances() (balances []Balance, err error) {
	r, err := b.client.do("GET", "account/getbalances", "", true)
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
	err = json.Unmarshal(response.Result, &balances)
	return
}

// Getbalance is used to retrieve the balance from your account for a specific currency.
// currency: a string literal for the currency (ex: LTC)
func (b *Bittrex) GetBalance(currency string) (balance Balance, err error) {
	r, err := b.client.do("GET", "account/getbalance?currency="+strings.ToUpper(currency), "", true)
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
	err = json.Unmarshal(response.Result, &balance)
	return
}

// GetDepositAddress is sed to generate or retrieve an address for a specific currency.
// currency a string literal for the currency (ie. BTC)
func (b *Bittrex) GetDepositAddress(currency string) (address Address, err error) {
	r, err := b.client.do("GET", "account/getdepositaddress?currency="+strings.ToUpper(currency), "", true)
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
	err = json.Unmarshal(response.Result, &address)
	return
}

// Withdraw is used to withdraw funds from your account.
// address string the address where to send the funds.
// currency string literal for the currency (ie. BTC)
// quantity decimal.Decimal the quantity of coins to withdraw
func (b *Bittrex) Withdraw(address, currency string, quantity decimal.Decimal) (withdrawUuid string, err error) {
	r, err := b.client.do("GET", fmt.Sprintf("account/withdraw?currency=%s&quantity=%s&address=%s", strings.ToUpper(currency), quantity, address), "", true)
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
	withdrawUuid = u.Id
	return
}

// GetOrderHistory used to retrieve your order history.
// market string literal for the market (ie. BTC-LTC). If set to "all", will return for all market
func (b *Bittrex) GetOrderHistory(market string) (orders []Order, err error) {
	resource := "account/getorderhistory"
	if market != "all" {
		resource += "?market=" + market
	}
	r, err := b.client.do("GET", resource, "", true)
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
	err = json.Unmarshal(response.Result, &orders)
	return
}

// GetWithdrawalHistory is used to retrieve your withdrawal history
// currency string a string literal for the currency (ie. BTC). If set to "all", will return for all currencies
func (b *Bittrex) GetWithdrawalHistory(currency string) (withdrawals []Withdrawal, err error) {
	resource := "account/getwithdrawalhistory"
	if currency != "all" {
		resource += "?currency=" + currency
	}
	r, err := b.client.do("GET", resource, "", true)
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
	err = json.Unmarshal(response.Result, &withdrawals)
	return
}

// GetDepositHistory is used to retrieve your deposit history
// currency string a string literal for the currency (ie. BTC). If set to "all", will return for all currencies
func (b *Bittrex) GetDepositHistory(currency string) (deposits []Deposit, err error) {
	resource := "account/getdeposithistory"
	if currency != "all" {
		resource += "?currency=" + currency
	}
	r, err := b.client.do("GET", resource, "", true)
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
	err = json.Unmarshal(response.Result, &deposits)
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
