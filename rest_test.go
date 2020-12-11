package bittrex

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)


func init() {
	_ = godotenv.Load()
}

func TestPublicAPI(t *testing.T)  {
	bittrex := New(os.Getenv("API_PUBLIC"), os.Getenv("API_SECRET"))
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


	// This should throw ERR_ORDER_MISSING_PARAMETERS
	resCreateOrder, err := bittrex.CreateOrder(CreateOrderParams{
		MarketSymbol:  "CRW-BTC",
		Direction:     BUY,
		Type:          "",
		Quantity:      decimal.Decimal{},
		Ceiling:       decimal.Decimal{},
		Limit:         decimal.Decimal{},
		TimeInForce:   "",
		ClientOrderID: "",
		UseAwards:     "",
	})
	assert.Equal(t, ERR_ORDER_MISSING_PARAMETERS, err)
	assert.Equal(t, OrderV3{}, resCreateOrder)

	resBalances, err := bittrex.GetBalances()
	assert.Nil(t, err)
	assert.NotNil(t, resBalances)
	/* for _, b := range resBalances {
		fmt.Println(b.CurrencySymbol, " - ", b.Total)
	}*/

	resBalance, err := bittrex.GetBalance("BTC")
	assert.Nil(t, err)
	assert.NotNil(t, resBalance)

	resAddress, err := bittrex.GetDepositAddress("BTC")
	assert.Nil(t, err)
	assert.NotNil(t, resAddress)
	fmt.Println(resAddress)
	// assert.Equal(t, "3", resAddress.CryptoAddress[0])


	resWithdrawalHistory, err := bittrex.GetOpenWithdrawals("", ALL)
	assert.Nil(t, err)
	assert.NotNil(t, resWithdrawalHistory)

	resClosedWithdrawalHistory, err := bittrex.GetClosedWithdrawals("", ALL)
	assert.Nil(t, err)
	assert.NotNil(t, resClosedWithdrawalHistory)
	// fmt.Println(resClosedWithdrawalHistory)

	resWithdrawalId, err := bittrex.GetWithdrawalByTxId("0xf962f9392673084411efc9ab7c83b18e1b6b6f34b0236af184cbf7001793862b")
	assert.Nil(t, err)
	assert.NotNil(t, resWithdrawalId)
	fmt.Println(resWithdrawalId)
}

/*func TestDeposits(t *testing.T)  {
	bittrex := New(os.Getenv("API_PUBLIC"), os.Getenv("API_SECRET"))
	info, err := bittrex.GetDepositHistory("9662cf066d792f32c19ab91c7b2959be7ed7ec6eac9f96bd6ffc00490a52d696")
	assert.Nil(t, err)
	fmt.Println(info)
}*/

/*func TestWithdraw(t *testing.T)  {
	bittrex := New(os.Getenv("API_PUBLIC"), os.Getenv("API_SECRET"))
	info, err := bittrex.Withdraw("CRWauZnzPf3oNDntLGhPZNbDRZWrPjJmrRS9", "CRW", decimal.NewFromFloat(5.0), "test")
	assert.Nil(t, err)
	fmt.Println(info)
}*/
