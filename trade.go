package bittrex

// Used in getmarkethistory
type Trade struct {
	OrderUuid string  `json:"OrderUuid"`
	TimeStamp string  `json:"TimeStamp"`
	Quantity  float64 `json:"Quantity"`
	Price     float64 `json:"Price"`
	Total     float64 `json:"Total"`
}
