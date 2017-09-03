package bittrex

import (
	"encoding/json"
	"fmt"

	"github.com/thebotguys/signalr"
)

type orderUpdate struct {
	Orderb
	Type int
}

type fill struct {
	Orderb
	OrderType string
	Timestamp jTime
}

type exchangeState struct {
	MarketName string
	Nounce     int
	Buys       []orderUpdate
	Sells      []orderUpdate
	Fills      []fill
}

func (b *Bittrex) SubscribeOrderBook(market string, stop <-chan bool) error {
	var lastNonce int
	var states []exchangeState
	client := signalr.NewWebsocketClient()
	client.OnClientMethod = func(hub string, method string, messages []json.RawMessage) {
		println(hub, method)
		if method != "updateExchangeState" {
			return
		}
		isFirst := lastNonce == 0 && len(states) == 0
		for _, msg := range messages {
			var st exchangeState
			if err := json.Unmarshal(msg, &st); err != nil {
				fmt.Println(err)
				continue
			}
			if st.MarketName != market {
				continue
			}
			if lastNonce == 0 {
				states = append(states, st)
				continue
			}
			fmt.Println(st)
		}
		if isFirst && len(states) > 0 {
			// we've got first exchange update. we can now request the entire state.
		}
	}
	if err := client.Connect("https", WS_BASE, []string{WS_HUB}); err != nil {
		return err
	}
	defer client.Close()
	_, err := client.CallHub(WS_HUB, "SubscribeToExchangeDeltas", market)
	if err != nil {
		return err
	}
	msg, err := client.CallHub(WS_HUB, "QueryExchangeState", market)
	if err != nil {
		return err
	}
	var st exchangeState
	if err := json.Unmarshal(msg, &st); err != nil {
		fmt.Println(err)
		return err
	}
	<-stop
	return nil
}
