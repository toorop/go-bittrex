package main

import (
	"fmt"
	"github.com/toorop/go-bittrex"
)

const (
	API_KEY = ""
)

func main() {
	// Bittrex client
	bittrex := bittrex.New(API_KEY)

	// Get markets
	markets, err := bittrex.GetMarkets()
	fmt.Println(err, markets)

	//
}
