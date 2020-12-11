package bittrex

import (
	"github.com/shopspring/decimal"
	"time"
)

type Balance struct {
	Currency      string          `json:"Currency"`
	Balance       decimal.Decimal `json:"Balance"`
	Available     decimal.Decimal `json:"Available"`
	Pending       decimal.Decimal `json:"Pending"`
	CryptoAddress string          `json:"CryptoAddress"`
	Requested     bool            `json:"Requested"`
	Uuid          string          `json:"Uuid"`
}

type BalanceV3 struct {
	CurrencySymbol string `json:"currencySymbol"`
	Total          decimal.Decimal `json:"total"`
	Available      decimal.Decimal `json:"available"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
