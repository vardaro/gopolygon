package main

import (
	"fmt"
	"os"

	"github.com/vardaro/gopolygon/models"
	"github.com/vardaro/gopolygon/stocks"
)

func main() {
	apikey := os.Getenv("POLYGON_API_KEY")
	if apikey == "" {
		fmt.Println("Cant find API Key")
		return
	}
	client := stocks.NewClient(apikey)

	// Query the last month of data
	query := &models.AggregatesQuery{
		Symbol:     "AAPL",
		Multiplier: 1,
		Timespan:   "day",
		From:       "2020-10-01",
		To:         "2020-11-01",
	}

	resp, err := client.Aggregates(query)
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Fail")
	}

	for i := 0; i < resp.ResultsCount; i++ {
		fmt.Println(resp.Results[i])
	}

}
