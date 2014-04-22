package bittrex

import (
	"encoding/json"
	"errors"
	//"fmt"
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
