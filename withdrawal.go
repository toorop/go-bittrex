package bittrex

import (
	"github.com/shopspring/decimal"
	"time"
)

type Withdrawal struct {
	PaymentUuid    string          `json:"PaymentUuid"`
	Currency       string          `json:"Currency"`
	Amount         decimal.Decimal `json:"Amount"`
	Address        string          `json:"Address"`
	Opened         jTime           `json:"Opened"`
	Authorized     bool            `json:"Authorized"`
	PendingPayment bool            `json:"PendingPayment"`
	TxCost         decimal.Decimal `json:"TxCost"`
	TxId           string          `json:"TxId"`
	Canceled       bool            `json:"Canceled"`
}

type WithdrawalParams struct {
	CurrencySymbol   string `json:"currencySymbol"`
	Quantity         string `json:"quantity"`
	CryptoAddress    string `json:"cryptoAddress"`
	CryptoAddressTag string `json:"cryptoAddressTag"`
}

type WithdrawalV3 struct {
	ID               string           `json:"id"`
	CurrencySymbol   string           `json:"currencySymbol"`
	Quantity         decimal.Decimal  `json:"quantity"`
	CryptoAddress    string           `json:"cryptoAddress"`
	CryptoAddressTag string           `json:"cryptoAddressTag"`
	TxCost           decimal.Decimal  `json:"txCost"`
	TxID             string           `json:"txId"`
	Status           WithdrawalStatus `json:"status"`
	CreatedAt        time.Time        `json:"createdAt"`
	CompletedAt      time.Time        `json:"completedAt"`
}

type WithdrawalHistoryParams struct {
	Status         string `url:"status,omitempty"`
	CurrencySymbol string `url:"currencySymbol,omitempty"`
}
