package bittrex

import "github.com/shopspring/decimal"

type Currency struct {
	Currency        string          `json:"Currency"`
	CurrencyLong    string          `json:"CurrencyLong"`
	MinConfirmation int             `json:"MinConfirmation"`
	TxFee           decimal.Decimal `json:"TxFee"`
	IsActive        bool            `json:"IsActive"`
	CoinType        string          `json:"CoinType"`
	BaseAddress     string          `json:"BaseAddress"`
	Notice          string          `json:"Notice"`
}
