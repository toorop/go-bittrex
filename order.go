package bittrex

type Order struct {
	OrderUuid         string  `json:"OrderUuid"`
	Exchange          string  `json:"Exchange"`
	TimeStamp         string  `json:"TimeStamp"`
	OrderType         string  `json:"OrderType"`
	Limit             float64 `json:"Limit"`
	Quantity          float64 `json:"Quantity"`
	QuantityRemaining float64 `json:"QuantityRemaining"`
	Commission        float64 `json:"Commission"`
	Price             float64 `json:"Price"`
	PricePerUnit      float64 `json:"PricePerUnit"`
}
