package bittrex

import "github.com/shopspring/decimal"

type Market struct {
	MarketCurrency     string  `json:"MarketCurrency"`
	BaseCurrency       string  `json:"BaseCurrency"`
	MarketCurrencyLong string  `json:"MarketCurrencyLong"`
	BaseCurrencyLong   string  `json:"BaseCurrencyLong"`
	MinTradeSize       decimal.Decimal `json:"MinTradeSize"`
	MarketName         string  `json:"MarketName"`
	IsActive           bool    `json:"IsActive"`
	IsRestricted       bool    `json:"IsRestricted"`
	Notice             string  `json:"Notice"`
	IsSponsored        bool    `json:"IsSponsored"`
	LogoUrl            string  `json:"LogoUrl"`
	Created            string  `json:"Created"`
}
