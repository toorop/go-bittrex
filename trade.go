package bittrex

// Used in getmarkethistory
type Trade struct {
	OrderUuid int64   `json:"Id"`
	Timestamp jTime   `json:"TimeStamp"`
	Quantity  float64 `json:"Quantity"`
	Price     float64 `json:"Price"`
	Total     float64 `json:"Total"`
	FillType  string  `json:"FillType"`
	OrderType string  `json:"OrderType"`
}
