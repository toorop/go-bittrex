package main

import (
	"fmt"
	"github.com/toorop/go-bittrex"
	//"time"
)

const (
	API_KEY = ""
)

func main() {
	// Bittrex client
	bittrex := bittrex.New(API_KEY)

	// Get markets
	//markets, err := bittrex.GetMarkets()
	//fmt.Println(err, markets)

	// Get Ticker (BTC-VTC)
	//ticker, err := bittrex.GetTicker("BTC-VTC")
	//fmt.Println(err, ticker)

	// Get market summaries
	//marketSummaries, err := bittrex.GetMarketSummaries()
	//fmt.Println(err, marketSummaries)

	// Get orders book
	/*orderBook, err := bittrex.GetOrderBook("BTC-QBC", "both", 100)
	fmt.Println(err, orderBook)
	*/

	// Market history
	marketHistory, err := bittrex.GetMarketHistory("BTC-QBC", 100)
	//fmt.Println(err, marketHistory)
	for _, trade := range marketHistory {
		fmt.Println(err, trade.Timestamp.String(), trade.Price)
	}

	// Account

	// Get balances
	/*balances, err := bittrex.GetBalances()
	fmt.Println(err, balances)*/

	// Get balance
	/*balance, err := bittrex.GetBalance("QBC")
	fmt.Println(err, balance)*/

	// Get address
	/*address, err := bittrex.GetDepositAddress("QBCT")
	fmt.Println(err, address)*/

	// WithDraw
	/*whitdrawUuid, err := bittrex.Withdraw("QYQeWgSnxwtTuW744z7Bs1xsgszWaFueQc", "QBC", 1.1)
	fmt.Println(err, whitdrawUuid)*/

	// Get order history
	/*orderHistory, err := bittrex.GetOrderHistory("QBC", 10)
	fmt.Println(err, orderHistory)*/

	// Get getwithdrawal history
	/*withdrawalHistory, err := bittrex.GetWithdrawalHistory("all", 0)
	fmt.Println(err, withdrawalHistory)*/

}
