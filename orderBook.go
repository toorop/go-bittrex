package bittrex

type OrderBook struct {
	Buy  []Order `json:"buy"`
	Sell []Order `json:"sell"`
}

type Order struct {
	Quantity float64 `json:"Quantity"`
	Rate     float64 `json:"Rate"`
}
