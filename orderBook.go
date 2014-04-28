package bittrex

type OrderBook struct {
	Buy  []Orderb `json:"buy"`
	Sell []Orderb `json:"sell"`
}

type Orderb struct {
	Quantity float64 `json:"Quantity"`
	Rate     float64 `json:"Rate"`
}
