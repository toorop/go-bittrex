package bittrex

import (
	"encoding/json"
)

type jsonResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}
