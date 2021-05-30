package bittrex

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	Uuid              *string
	OrderUuid         string          `json:"OrderUuid"`
	Exchange          string          `json:"Exchange"`
	OrderType         string          `json:"OrderType"`
	Limit             decimal.Decimal `json:"Limit"`
	Quantity          decimal.Decimal `json:"Quantity"`
	QuantityRemaining decimal.Decimal `json:"QuantityRemaining"`
	Price             decimal.Decimal `json:"Price"`
	PricePerUnit      decimal.Decimal `json:"PricePerUnit"`
	CommissionPaid    decimal.Decimal
	Opened            jTime
	Closed            *jTime
	CancelInitiated   bool
	ImmediateOrCancel bool
	IsConditional     bool
	Condition         string
	ConditionTarget   decimal.Decimal
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
	Opened                     jTime
	Closed                     *jTime
	IsOpen                     bool
	Sentinel                   string
	CancelInitiated            bool
	ImmediateOrCancel          bool
	IsConditional              bool
	Condition                  string
	ConditionTarget            decimal.Decimal
}

type CreateOrderParams struct {
	MarketSymbol string          `json:"marketSymbol"`
	Direction    OrderDirection  `json:"direction"`
	Type         OrderType       `json:"type"`
	Quantity     decimal.Decimal `json:"quantity"`
	TimeInForce  TimeInForce     `json:"timeInForce"`

	Ceiling       float64 `json:"ceiling,omitempty"`
	Limit         float64 `json:"limit,omitempty"`
	ClientOrderID string  `json:"clientOrderId,omitempty"`
	UseAwards     string  `json:"useAwards,omitempty"`
}

type OrderV3 struct {
	ID            string          `json:"id"`
	MarketSymbol  string          `json:"marketSymbol"`
	Direction     string          `json:"direction"`
	Type          string          `json:"type"`
	Quantity      decimal.Decimal `json:"quantity"`
	Limit         decimal.Decimal `json:"limit"`
	Ceiling       decimal.Decimal `json:"ceiling"`
	TimeInForce   string          `json:"timeInForce"`
	ClientOrderID string          `json:"clientOrderId"`
	FillQuantity  decimal.Decimal `json:"fillQuantity"`
	Commission    decimal.Decimal `json:"commission"`
	Proceeds      decimal.Decimal `json:"proceeds"`
	Status        string          `json:"status"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	ClosedAt      time.Time       `json:"closedAt"`
	OrderToCancel OrderData       `json:"orderToCancel"`
}

type OrderData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
