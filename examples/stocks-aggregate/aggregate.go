package main

import (
	"fmt"
	"os"
	"time"

	"github.com/vardaro/gopolygon"
)

func main() {
	apikey := os.Getenv(gopolygon.EnvAPIKey)
	if apikey == "" {
		fmt.Println("Cant find API Key")
		return
	}
	client := gopolygon.NewClient(apikey)

	// Print price data from the last month
	now := time.Now()
	lastmonth := now.AddDate(0, -1, 0)
	unadjusted := true
	resp, err := client.Aggregates("AAPL", 1, "day", &lastmonth, &now, &unadjusted)
	if err != nil {
		fmt.Println("Fail")
	}

	for i := 0; i < resp.ResultsCount; i++ {
		fmt.Println(resp.Results[i])
	}

}
