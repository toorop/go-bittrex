package bittrex

import "github.com/shopspring/decimal"

type Ticker struct {
	Bid  decimal.Decimal `json:"Bid"`
	Ask  decimal.Decimal `json:"Ask"`
	Last decimal.Decimal `json:"Last"`
}
