package bittrex

import "github.com/shopspring/decimal"

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
