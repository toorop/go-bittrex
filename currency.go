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

type CurrencyV3 struct {
	Symbol                   string          `json:"symbol"`
	Name                     string          `json:"name"`
	CoinType                 string          `json:"coinType"`
	Status                   string          `json:"status"`
	MinConfirmations         int             `json:"minConfirmations"`
	Notice                   string          `json:"notice"`
	TxFee                    decimal.Decimal `json:"txFee"`
	LogoURL                  string          `json:"logoUrl,omitempty"`
	ProhibitedIn             []interface{}   `json:"prohibitedIn"`
	BaseAddress              string          `json:"baseAddress,omitempty"`
	AssociatedTermsOfService []interface{}   `json:"associatedTermsOfService"`
}
