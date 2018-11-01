package bittrex

import "github.com/shopspring/decimal"

type Candle struct {
	TimeStamp  CandleTime      `json:"T"`
	Open       decimal.Decimal `json:"O"`
	Close      decimal.Decimal `json:"C"`
	High       decimal.Decimal `json:"H"`
	Low        decimal.Decimal `json:"L"`
	Volume     decimal.Decimal `json:"V"`
	BaseVolume decimal.Decimal `json:"BV"`
}

type NewCandles struct {
	Ticks []Candle `json:"ticks"`
}
