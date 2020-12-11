package bittrex

import (
	"github.com/shopspring/decimal"
	"time"
)

// Used in getmarkethistory
type Trade struct {
	OrderUuid int64           `json:"Id"`
	Timestamp jTime           `json:"TimeStamp"`
	Quantity  decimal.Decimal `json:"Quantity"`
	Price     decimal.Decimal `json:"Price"`
	Total     decimal.Decimal `json:"Total"`
	FillType  string          `json:"FillType"`
	OrderType string          `json:"OrderType"`
}

type TradeV3 struct {
	ID         string `json:"id"`
	ExecutedAt time.Time `json:"executedAt"`
	Quantity   string `json:"quantity"`
	Rate       string `json:"rate"`
	TakerSide  string `json:"takerSide"`
}