package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var currency = "USD"

type Gold struct {
	Prices []Price `json:"items"`
	Client *http.Client
}

type Price struct {
	Currency      string    `json:"currency"`
	Price         float64   `json:"xauPrice"`
	Change        float64   `json:"chgXau"`
	PreviousClose float64   `json:"xauClose"`
	Time          time.Time `json:"-"`
}

func (g *Gold) GetPrices() (*Price, error) {
	if g.Client == nil {
		g.Client = &http.Client{}
	}

	client := g.Client
	url := fmt.Sprintf("https://data-asg.goldprice.org/dbXRates/%s", currency)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error while fetching gold prices")
		return nil, fmt.Errorf("failed to get a response: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error while reading body")
		return nil, fmt.Errorf("failed to read a body: %w", err)
	}

	gold := Gold{}

	err = json.Unmarshal(body, &gold)
	if err != nil {
		log.Println("error while unmarshal")
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	var prev, current, change float64
	prev, current, change = gold.Prices[0].PreviousClose, gold.Prices[0].Price, gold.Prices[0].Change

	var currentInfo = Price{
		Currency:      currency,
		Price:         current,
		Change:        change,
		PreviousClose: prev,
		Time:          time.Now(),
	}

	return &currentInfo, nil
}
