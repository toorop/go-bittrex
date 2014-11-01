package bittrex

type Candle struct {
	TimeStamp CandleTime `json:"T"`
	Open      float64    `json:"O"`
	Close     float64    `json:"C"`
	High      float64    `json:"H"`
	Low       float64    `json:"L"`
	Volume    float64    `json:"V"`
}

type NewCandles struct {
	Ticks []Candle `json:"ticks"`
}
