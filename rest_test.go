package bittrex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublicAPI(t *testing.T)  {
	bittrex := New("", "")
	_, err := bittrex.GetCurrencies()
	assert.Nil(t, err)

	_, err = bittrex.GetCurrency("DGB")
	assert.Nil(t, err)

	_, err = bittrex.GetMarkets()
	assert.Nil(t, err)

	_, err = bittrex.GetTicker("")
	assert.Nil(t, err)

	_, err = bittrex.GetTicker("CRW-BTC")
	assert.Nil(t, err)

	_, err =  bittrex.GetMarketSummaries()
	assert.Nil(t, err)

	res, err :=  bittrex.GetMarketSummary("CRW-BTC")
	assert.Nil(t, err)
	fmt.Println(res)
}
