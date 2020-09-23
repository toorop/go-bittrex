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

	_, err =  bittrex.GetMarketSummary("CRW-BTC")
	assert.Nil(t, err)

	// TODO Verify this test cases

	resOrderBook, err :=  bittrex.GetOrderBook("CRW-BTC", 500, "buy")
	assert.Nil(t, err)
	assert.Nil(t, resOrderBook.Bid)
	assert.NotNil(t, resOrderBook.Ask)

	resOrderBook2, err :=  bittrex.GetOrderBook("CRW-BTC", 500, "sell")
	assert.Nil(t, err)
	assert.Nil(t, resOrderBook2.Ask)
	assert.NotNil(t, resOrderBook2.Bid)

	resOrderBook3, err :=  bittrex.GetOrderBook("CRW-BTC", 500, "")
	assert.Nil(t, err)
	assert.NotNil(t, resOrderBook3.Bid)
	assert.NotNil(t, resOrderBook3.Ask)

	resOrderBook4, err :=  bittrex.GetOrderBook("CRW-BTC", 500, "both")
	assert.Nil(t, err)
	assert.NotNil(t, resOrderBook4.Bid)
	assert.NotNil(t, resOrderBook4.Ask)

	resOrderBookSide, err :=  bittrex.GetOrderBookBuySell("CRW-BTC", 500, "buy")
	assert.Nil(t, err)
	assert.NotNil(t, resOrderBookSide)

	resTrades, err :=  bittrex.GetMarketHistory("CRW-BTC")
	assert.Nil(t, err)
	assert.NotNil(t, resTrades)
	fmt.Println(resTrades)
}
