package bittrex

import (
	"testing"
)

func TestBittrexSubscribeOrderBook(t *testing.T) {
	bt := New("", "")
	if err := bt.SubscribeOrderBook("BTC-LTC", nil); err != nil {
		t.Error(err)
	}
}
