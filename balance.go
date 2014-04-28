package bittrex

type Balance struct {
	Currency  string  `json:"Currency"`
	Balance   float64 `json:"Balance"`
	Available float64 `json:"Available"`
	Pending   float64 `json:"Pending"`
}
