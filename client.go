package bittrex

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) (c *client) {
	return &client{apiKey, &http.Client{}}
}

func (c *client) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	// Do the request in the background so we can check the timeout
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := c.httpClient.Do(req)
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		// Kill the connection on timeout so we don't leak sockets or goroutines
		//c.client.Transport.CancelRequest(req)
		return nil, errors.New("timeout on reading data from Bittrex API")
	}
}

func (c *client) do(method string, ressource string, payload string) (response []byte, err error) {
	connectTimer := time.NewTimer(DEFAULT_HTTPCLIENT_TIMEOUT * time.Second)

	query := fmt.Sprintf("%s/%s/%s", API_BASE, API_VERSION, ressource)
	req, err := http.NewRequest(method, query, strings.NewReader(payload))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")

	req.Header.Set("APIKEY", c.apiKey)
	resp, err := c.doTimeoutRequest(connectTimer, req)
	//fmt.Println(resp.Status, resp.StatusCode, err)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	//fmt.Println(fmt.Sprintf("reponse %s", response), err)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return response, err
}
