package main

import (
	"fmt"
	"os"

	gpoly "github.com/vardaro/gopolygon"
)

func main() {
	apikey := os.Getenv(gpoly.EnvAPIKey)
	if apikey == "" {
		fmt.Println("Cant find API Key")
		return
	}
	client := gpoly.NewClient(apikey)

	// Query the last month of data
	query := &gpoly.AggregatesQuery{
		Symbol:     "AAPL",
		Multiplier: 1,
		Timespan:   "day",
		From:       "2020-10-01",
		To:         "2020-11-01",
		Unadjusted: gpoly.Bool(true),
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
