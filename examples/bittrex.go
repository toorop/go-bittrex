package main

import (
	"fmt"
	"github.com/toorop/go-bittrex"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func main() {
	// Bittrex client
	bittrex := bittrex.New(API_KEY, API_SECRET)

	// Get Candle ( OHLCV )
	/*
		markets, err := bittrex.GetHisCandles("BTC-LTC", "hour")
		fmt.Println(markets, err)
	*/

	// Get markets
	/*
		markets, err := bittrex.GetMarkets()
		fmt.Println(err, markets)
	*/

	// Get Ticker (BTC-VTC)
	/*
		ticker, err := bittrex.GetTicker("BTC-DRK")
		fmt.Println(err, ticker)
	*/

	// Get Distribution (JBS)
	/*
		distribution, err := bittrex.GetDistribution("JBS")
		for _, balance := range distribution.Distribution {
			fmt.Println(balance.BalanceD)
		}
	*/

	// Get market summaries
	/*
		marketSummaries, err := bittrex.GetMarketSummaries()
		fmt.Println(err, marketSummaries)
	*/

	// Get market summary
	/*
		marketSummary, err := bittrex.GetMarketSummary("BTC-ETH")
		fmt.Println(err, marketSummary)
	*/

	// Get orders book
	/*
		orderBook, err := bittrex.GetOrderBook("BTC-DRK", "both", 100)
		fmt.Println(err, orderBook)
	*/

	// Get order book buy or sell side only
	/*
		orderb, err := bittrex.GetOrderBookBuySell("BTC-JBS", "buy", 100)
		fmt.Println(err, orderb)
	*/

	// Market history
	/*
		marketHistory, err := bittrex.GetMarketHistory("BTC-DRK", 100)
		for _, trade := range marketHistory {
			fmt.Println(err, trade.Timestamp.String(), trade.Quantity, trade.Price)
		}
	*/

	// Market

	// BuyLimit
	/*
		uuid, err := bittrex.BuyLimit("BTC-DOGE", 1000, 0.00000102)
		fmt.Println(err, uuid)
	*/

	// BuyMarket
	/*
		uuid, err := bittrex.BuyLimit("BTC-DOGE", 1000)
		fmt.Println(err, uuid)
	*/

	// Sell limit
	/*
		uuid, err := bittrex.SellLimit("BTC-DOGE", 1000, 0.00000115)
		fmt.Println(err, uuid)
	*/

	// Cancel Order
	/*
		err := bittrex.CancelOrder("e3b4b704-2aca-4b8c-8272-50fada7de474")
		fmt.Println(err)
	*/

	// Get open orders
	/*
		orders, err := bittrex.GetOpenOrders("BTC-DOGE")
		fmt.Println(err, orders)
	*/

	// Account
	// Get balances
	/*
		balances, err := bittrex.GetBalances()
		fmt.Println(err, balances)
	*/

	// Get balance
	/*
		balance, err := bittrex.GetBalance("DOGE")
		fmt.Println(err, balance)
	*/

	// Get address
	/*
		address, err := bittrex.GetDepositAddress("QBC")
		fmt.Println(err, address)
	*/

	// WithDraw
	/*
		whitdrawUuid, err := bittrex.Withdraw("QYQeWgSnxwtTuW744z7Bs1xsgszWaFueQc", "QBC", 1.1)
		fmt.Println(err, whitdrawUuid)
	*/

	// Get order history
	/*
		orderHistory, err := bittrex.GetOrderHistory("BTC-DOGE", 10)
		fmt.Println(err, orderHistory)
	*/

	// Get getwithdrawal history
	/*
		withdrawalHistory, err := bittrex.GetWithdrawalHistory("all", 0)
		fmt.Println(err, withdrawalHistory)
	*/

	// Get deposit history
	/*
		deposits, err := bittrex.GetDepositHistory("all", 0)
		fmt.Println(err, deposits)
	*/

}
