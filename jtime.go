package bittrex

import (
	"encoding/json"
	//"fmt"
	"time"
)

const TIME_FORMAT = "2006-01-02T15:04:05"

type jTime struct {
	time.Time
}

func (jt *jTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(TIME_FORMAT, s)
	if err != nil {
		return err
	}
	jt.Time = t
	return nil
}

func (jt jTime) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&jt.Time).Format(TIME_FORMAT))
}
