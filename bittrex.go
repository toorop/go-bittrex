package bittrex

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	API_BASE                   = "https://bittrex.com/api"
	API_VERSION                = "v1.1"
	DEFAULT_HTTPCLIENT_TIMEOUT = 30
)

type bittrex struct {
	client *client
}

func New(apiKey string) *bittrex {
	client := NewClient(apiKey)
	return &bittrex{client}
}

// GetMarkets is used to get the open and available trading markets at Bittrex along with other meta data.
func (b *bittrex) GetMarkets() (markets []Market, err error) {
	r, err := b.client.do("GET", "/public/getmarkets", "")
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if !response.Success {
		err = errors.New(response.Message)
		return
	}
	json.Unmarshal(response.Result, &markets)
	return
}

// GetTicker is used to get the current tick values for a market.
func (b *bittrex) GetTicker(market string) (ticker Ticker, err error) {
	r, err := b.client.do("GET", "/public/getticker?market="+strings.ToUpper(market), "")
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if !response.Success {
		err = errors.New(response.Message)
		return
	}
	json.Unmarshal(response.Result, &ticker)
	return
}

// GetMarketSummaries is used to get the last 24 hour summary of all active exchanges
func (b *bittrex) GetMarketSummaries() (marketSummaries []MarketSummary, err error) {
	r, err := b.client.do("GET", "/public/getmarketsummaries", "")
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if !response.Success {
		err = errors.New(response.Message)
		return
	}
	json.Unmarshal(response.Result, &marketSummaries)
	return
}

// GetOrderBook is used to get retrieve the orderbook for a given market
// market: a string literal for the market (ex: BTC-LTC)
// cat: buy, sell or both to identify the type of orderbook to return.
// depth: how deep of an order book to retrieve. Max is 100

func (b *bittrex) GetOrderBook(market, cat string, depth int) (orderBook OrderBook, err error) {
	if cat != "buy" && cat != "sell" && cat != "both" {
		cat = "both"
	}
	if depth > 100 {
		depth = 100
	}

	req := fmt.Sprintf("/public/getorderbook?market=%s&type=%s&depth=%d", strings.ToUpper(market), cat, depth)
	r, err := b.client.do("GET", req, "")
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if !response.Success {
		err = errors.New(response.Message)
		return
	}
	json.Unmarshal(response.Result, &orderBook)
	return
}
