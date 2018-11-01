package bittrex

import (
	"fmt"
	"time"
)

var CANDLE_INTERVALS = map[string]bool{
	"oneMin":    true,
	"fiveMin":   true,
	"thirtyMin": true,
	"hour":      true,
	"day":       true,
}

type CandleTime struct {
	time.Time
}

func (t *CandleTime) UnmarshalJSON(b []byte) error {
	if len(b) < 2 {
		return fmt.Errorf("could not parse time %s", string(b))
	}
	// trim enclosing ""
	result, err := time.Parse("2006-01-02T15:04:05", string(b[1:len(b)-1]))
	if err != nil {
		return fmt.Errorf("could not parse time: %v", err)
	}
	t.Time = result
	return nil
}
