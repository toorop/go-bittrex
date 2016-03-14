package bittrex

type Distribution struct {
	Distribution   []BalanceD `json:"Distribution"`
	Balances       float64    `json:"Balances"`
	AverageBalance float64    `json:"AverageBalance"`
}

type BalanceD struct {
	BalanceD float64 `json:"Balance"`
}
