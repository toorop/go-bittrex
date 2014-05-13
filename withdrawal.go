package bittrex

type Withdrawal struct {
	PaymentUuid    string  `json:"PaymentUuid"`
	Currency       string  `json:"Currency"`
	Amount         float64 `json:"Amount"`
	Address        string  `json:"Address"`
	Opened         jTime   `json:"Opened"`
	Authorized     bool    `json:"Authorized"`
	PendingPayment bool    `json:"PendingPayment"`
	TxCost         float64 `json:"TxCost"`
	TxId           string  `json:"TxId"`
	Canceled       bool    `json:"Canceled"`
}
