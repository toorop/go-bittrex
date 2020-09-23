package bittrex

import "github.com/shopspring/decimal"

type OrderBook struct {
	Buy  []Orderb `json:"buy"`
	Sell []Orderb `json:"sell"`
}

type Orderb struct {
	Quantity decimal.Decimal `json:"Quantity"`
	Rate     decimal.Decimal `json:"Rate"`
}

type OrderBookV3 struct {
	Bid []OrderbV3 `json:"bid"`
	Ask []OrderbV3 `json:"ask"`
}

type OrderbV3 struct {
	Quantity decimal.Decimal `json:"quantity"`
	Rate     decimal.Decimal `json:"rate"`
}