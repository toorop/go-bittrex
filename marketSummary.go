package bittrex

type MarketSummary struct {
	MarketName string  `json:"MarketName"`
	High       float64 `json:"High"`
	Low        float64 `json:"Low"`
	Volume     float64 `json:"Volume"`
	Last       float64 `json:"Last"`
	BaseVolume float64 `json:"BaseVolume"`
	TimeStamp  string  `json:"TimeStamp"`
}
