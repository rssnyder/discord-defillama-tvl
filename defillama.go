package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	DefiLlamaTVLURL = "https://api.llama.fi/tvl/%s"
)

func GetTVL(protocol string) (result float64, err error) {

	req, err := http.NewRequest("GET", fmt.Sprintf(DefiLlamaTVLURL, protocol), nil)
	if err != nil {
		return result, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	results, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	return strconv.ParseFloat(string(results), 32)
}
