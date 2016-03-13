package bittrex

type Order struct {
	OrderUuid         string  `json:"OrderUuid"`
	Exchange          string  `json:"Exchange"`
	TimeStamp         jTime   `json:"TimeStamp"`
	OrderType         string  `json:"OrderType"`
	Limit             float64 `json:"Limit"`
	Quantity          float64 `json:"Quantity"`
	QuantityRemaining float64 `json:"QuantityRemaining"`
	Commission        float64 `json:"Commission"`
	Price             float64 `json:"Price"`
	PricePerUnit      float64 `json:"PricePerUnit"`
}

// For getorder
type Order2 struct {
	AccountId                  string
	OrderUuid                  string `json:"OrderUuid"`
	Exchange                   string `json:"Exchange"`
	Type                       string
	Quantity                   float64 `json:"Quantity"`
	QuantityRemaining          float64 `json:"QuantityRemaining"`
	Limit                      float64 `json:"Limit"`
	Reserved                   float64
	ReserveRemaining           float64
	CommissionReserved         float64
	CommissionReserveRemaining float64
	CommissionPaid             float64
	Price                      float64 `json:"Price"`
	PricePerUnit               float64 `json:"PricePerUnit"`
	Opened                     string
	Closed                     string
	IsOpen                     bool
	Sentinel                   string
	CancelInitiated            bool
	ImmediateOrCancel          bool
	IsConditional              bool
	Condition                  string
	ConditionTarget            string
}
