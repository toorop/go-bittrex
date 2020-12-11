package bittrex

import "github.com/shopspring/decimal"

type Ticker struct {
	Bid  decimal.Decimal `json:"Bid"`
	Ask  decimal.Decimal `json:"Ask"`
	Last decimal.Decimal `json:"Last"`
}

type TickerV3 struct {
	Symbol        string `json:"symbol"`
	LastTradeRate decimal.Decimal `json:"lastTradeRate"`
	BidRate       decimal.Decimal `json:"bidRate"`
	AskRate       decimal.Decimal `json:"askRate"`
}
