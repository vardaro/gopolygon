package main

import (
	"fmt"
	"os"
	"time"

	gpoly "github.com/vardaro/gopolygon"
)

func main() {
	apikey := os.Getenv(gpoly.EnvAPIKey)
	if apikey == "" {
		fmt.Println("Cant find API Key")
		return
	}
	client := gpoly.NewClient(apikey)

	// Print price data from the last month
	now := time.Now()
	lastmonth := now.AddDate(0, -1, 0)

	query := &gpoly.AggregatesQuery{
		Symbol:     "AAPL",
		Multiplier: 1,
		Timespan:   "day",
		From:       lastmonth,
		To:         now,
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
