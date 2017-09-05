package bittrex

import (
	"encoding/json"
	"fmt"

	"github.com/thebotguys/signalr"
)

type OrderUpdate struct {
	Orderb
	Type int
}

type Fill struct {
	Orderb
	OrderType string
	Timestamp jTime
}

type ExchangeState struct {
	MarketName string
	Nounce     int
	Buys       []OrderUpdate
	Sells      []OrderUpdate
	Fills      []Fill
	Initial    bool
}

func (b *Bittrex) SubscribeExchangeUpdate(market string, dataCh chan<- ExchangeState, stop <-chan bool) error {
	client := signalr.NewWebsocketClient()
	sendUpdate := func(st ExchangeState) {
		select {
		case dataCh <- st:
		default:
		}
	}
	client.OnClientMethod = func(hub string, method string, messages []json.RawMessage) {
		if hub != "CoreHub" || method != "updateExchangeState" {
			return
		}
		for _, msg := range messages {
			var st ExchangeState
			if err := json.Unmarshal(msg, &st); err != nil {
				fmt.Println(err)
				continue
			}
			if st.MarketName != market {
				continue
			}
			sendUpdate(st)
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
	var st ExchangeState
	if err := json.Unmarshal(msg, &st); err != nil {
		return err
	}
	st.Initial = true
	sendUpdate(st)
	<-stop
	return nil
}
