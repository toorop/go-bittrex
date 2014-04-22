package bittrex

type Ticker struct {
	Bid  float64 `json:"Bid"`
	Ask  float64 `json:"Ask"`
	Last float64 `json:"Last"`
}
