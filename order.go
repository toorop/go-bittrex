package bittrex

import "github.com/shopspring/decimal"

type Order struct {
	OrderUuid         string          `json:"OrderUuid"`
	Exchange          string          `json:"Exchange"`
	TimeStamp         jTime           `json:"TimeStamp"`
	OrderType         string          `json:"OrderType"`
	Limit             decimal.Decimal `json:"Limit"`
	Quantity          decimal.Decimal `json:"Quantity"`
	QuantityRemaining decimal.Decimal `json:"QuantityRemaining"`
	Commission        decimal.Decimal `json:"Commission"`
	Price             decimal.Decimal `json:"Price"`
	PricePerUnit      decimal.Decimal `json:"PricePerUnit"`
}

// For getorder
type Order2 struct {
	AccountId                  string
	OrderUuid                  string `json:"OrderUuid"`
	Exchange                   string `json:"Exchange"`
	Type                       string
	Quantity                   decimal.Decimal `json:"Quantity"`
	QuantityRemaining          decimal.Decimal `json:"QuantityRemaining"`
	Limit                      decimal.Decimal `json:"Limit"`
	Reserved                   decimal.Decimal
	ReserveRemaining           decimal.Decimal
	CommissionReserved         decimal.Decimal
	CommissionReserveRemaining decimal.Decimal
	CommissionPaid             decimal.Decimal
	Price                      decimal.Decimal `json:"Price"`
	PricePerUnit               decimal.Decimal `json:"PricePerUnit"`
	Opened                     string
	Closed                     string
	IsOpen                     bool
	Sentinel                   string
	CancelInitiated            bool
	ImmediateOrCancel          bool
	IsConditional              bool
	Condition                  string
	ConditionTarget            decimal.Decimal
}
