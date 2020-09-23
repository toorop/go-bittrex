package bittrex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublicAPI(t *testing.T)  {
	bittrex := New("", "")
	_, err := bittrex.GetCurrencies()
	assert.Nil(t, err)

	_, err = bittrex.GetMarkets()
	assert.Nil(t, err)

	
}
