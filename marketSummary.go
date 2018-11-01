package bittrex

import "github.com/shopspring/decimal"

type MarketSummary struct {
	MarketName     string          `json:"MarketName"`
	High           decimal.Decimal `json:"High"`
	Low            decimal.Decimal `json:"Low"`
	Ask            decimal.Decimal `json:"Ask"`
	Bid            decimal.Decimal `json:"Bid"`
	OpenBuyOrders  int             `json:"OpenBuyOrders"`
	OpenSellOrders int             `json:"OpenSellOrders"`
	Volume         decimal.Decimal `json:"Volume"`
	Last           decimal.Decimal `json:"Last"`
	BaseVolume     decimal.Decimal `json:"BaseVolume"`
	PrevDay        decimal.Decimal `json:"PrevDay"`
	TimeStamp      string          `json:"TimeStamp"`
}
