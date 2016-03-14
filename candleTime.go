package bittrex

import (
	"strconv"
	"time"
)

var CANDLE_INTERVALS = map[string]bool{
	"tenmin": true,
	"hour":   true,
	"day":    true,
}

type CandleTime struct {
	time.Time
}

func (t *CandleTime) UnmarshalJSON(b []byte) error {
	result, err := strconv.ParseInt(string(b)[8:18], 10, 64)

	if err != nil {
		return err
	}

	t.Time = time.Unix(result, 0)

	return nil
}
