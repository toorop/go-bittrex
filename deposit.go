package bittrex

type Deposit struct {
	Id            int64   `json:"Id"`
	Amount        float64 `json:"Amount"`
	Currency      string  `json:"Currency"`
	Confirmations int     `json:"Confirmations"`
	LastUpdated   jTime   `json:"LastUpdated"`
	TxId          string  `json:"TxId"`
	CryptoAddress string  `json:"CryptoAddress"`
}
