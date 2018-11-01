go-bittrex [![GoDoc](https://godoc.org/github.com/toorop/go-bittrex?status.svg)](https://godoc.org/github.com/toorop/go-bittrex)
==========

go-bittrex is an implementation of the Bittrex API (public and private) in Golang.

This version implement V1.1 Bittrex API and the new HMAC authentification.

## Import
	import "github.com/toorop/go-bittrex"
	
## Usage

In order to use the client with go's default http client settings you can do:

~~~ go
package main

import (
	"fmt"
	"github.com/toorop/go-bittrex"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// Bittrex client
	bittrex := bittrex.New(API_KEY, API_SECRET)

	// Get markets
	markets, err := bittrex.GetMarkets()
	fmt.Println(err, markets)
}
~~~

In order to use custom settings for the http client do:

~~~ go
package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/toorop/go-bittrex"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	// Bittrex client
	bittrex := bittrex.NewWithCustomHttpClient(API_KEY, API_SECRET, httpClient)

	// Get markets
	markets, err := bittrex.GetMarkets()
	fmt.Println(err, markets)
}
~~~

See ["Examples" folder for more... examples](https://github.com/toorop/go-bittrex/blob/master/examples/bittrex.go)

## Documentation
[![GoDoc](https://godoc.org/github.com/toorop/go-bittrex?status.png)](https://godoc.org/github.com/toorop/go-bittrex)


## Stay tuned
[Follow me on Twitter](https://twitter.com/poroot)

Donate
------

![Donation QR](http://api.qrserver.com/v1/create-qr-code/?size=200x200&data=bitcoin:1HgpsmxV52eAjDcoNpVGpYEhGfgN7mM1JB%3Flabel%3Dtoorop)

[1HgpsmxV52eAjDcoNpVGpYEhGfgN7mM1JB](http://tinyurl.com/mccsoez)
