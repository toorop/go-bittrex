package bittrex

import "github.com/shopspring/decimal"

type Deposit struct {
	Id            int64           `json:"Id"`
	Amount        decimal.Decimal `json:"Amount"`
	Currency      string          `json:"Currency"`
	Confirmations int             `json:"Confirmations"`
	LastUpdated   jTime           `json:"LastUpdated"`
	TxId          string          `json:"TxId"`
	CryptoAddress string          `json:"CryptoAddress"`
}

type DepositV3 struct {
	ID               string          `json:"id"`
	CurrencySymbol   string          `json:"currencySymbol"`
	Quantity         decimal.Decimal `json:"quantity"`
	CryptoAddress    string          `json:"cryptoAddress"`
	CryptoAddressTag string          `json:"cryptoAddressTag"`
	TxID             string          `json:"txId"`
	Confirmations    int32          `json:"confirmations"`
	UpdatedAt        string          `json:"updatedAt"`
	CompletedAt      string          `json:"completedAt"`
	Status           string          `json:"status"`
	Source           string          `json:"source"`
}

type DepositHistoryParams struct {
	Status         string `url:"status,omitempty"`
	CurrencySymbol string `url:"currencySymbol,omitempty"`
}
