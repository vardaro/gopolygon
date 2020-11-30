# gopolygon
Go SDK for Polygon API.

## Installation
```go get github.com/vardaro/gopolygon```

## Example

Example querying the closing price of AAPL on 03/20/2020.

```golang
package main

import (
	"fmt"

	"github.com/vardaro/gopolygon/models"
	"github.com/vardaro/gopolygon/stocks"
)

func main() {
	client := stocks.NewClient("my_api_key")

	query := &models.DailyOpenCloseQuery{
		Symbol: "AAPL",
		Date:   "2020-03-20",
	}

	aapl, _ := client.DailyOpenClose(query)

	fmt.Println(aapl.Close)
}
```

The `models` package contains structs for each query and response to/from the Polygon API. In this example, a `DailyOpenCloseQuery` object is passed and a `DailyOpenCloseResponse` object is returned. 

Not required API params are pointers in their respective so they can be ommitted if the user does not want to use them.

## Supported Routes and Their Signatures

### Stocks

#### Daily Open/Close
```golang
func (c *Client) DailyOpenClose(opts *models.DailyOpenCloseQuery) (*models.DailyOpenCloseResponse, error)

type DailyOpenCloseQuery struct {
	Symbol string
	Date   string
}

type DailyOpenCloseResponse struct {
	Symbol     string  `json:"symbol"`
	From       string  `json:"from"`
	Open       float64 `json:"open"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Close      float64 `json:"close"`
	Volume     float64 `json:"volume"`
	PreMarket  float64 `json:"preMarket"`
	AfterHours float64 `json:"afterHours"`
}
```

#### Grouped Daily (Bars)
```golang
func (c *Client) GroupedDailyBars(opts *models.GroupedDailyBarsQuery) (*models.GroupedDailyBarsResponse, error)

type GroupedDailyBarsQuery struct {
	Date       string
	Unadjusted *bool
}

type GroupedDailyBarsResponse struct {
	Status       string `json:"status"`
	Adjusted     bool   `json:"adjusted"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	RequestID    string `json:"request_id"`
	Results      []Bar  `json:"results"`
}
```


#### Previous Close
```golang
func (c *Client) PreviousClose(opts *models.PreviousCloseQuery) (*models.PreviousCloseResponse, error)

type PreviousCloseQuery struct {
	Symbol     string
	Unadjusted *bool
}

type PreviousCloseResponse struct {
	Ticker       string `json:"ticker"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Bar  `json:"results"`
}
```

#### Aggregates (Bar)
```golang
func (c *Client) Aggregates(opts *models.AggregatesQuery) (*models.AggregatesResponse, error)

type AggregatesQuery struct {
	Symbol     string
	Multiplier int
	Timespan   string
	From       string
	To         string
	Unadjusted *bool
	Sort       *string
	Limit      *int
}

type AggregatesResponse struct {
	Ticker       string `json:"ticker"`
	Status       string `json:"status"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Bar  `json:"results"`
}
```

#### Reference, Crypto, Forex will be added in the future. 