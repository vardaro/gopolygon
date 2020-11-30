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

	fmt.Println(aapl.High)
}
```

The `models` package contains structs for each query and response to/from the Polygon API. In this example, a `DailyOpenCloseQuery` object is passed and a `DailyOpenCloseResponse` object is returned. 

Not required API params are pointers in their respective so they can be ommitted if the user does not want to use them.

## Supported Routes and Their Signatures

### Stocks
- Daily Open/Close
- - ```DailyOpenClose(opts *models.DailyOpenCloseQuery) (*models.DailyOpenCloseResponse, error)```


- Grouped Daily (Bars)
- - ```GroupedDailyBars(opts *models.GroupedDailyBarsQuery) (*models.GroupedDailyBarsResponse, error)```


- Previous Close
- - ```PreviousClose(opts *models.PreviousCloseQuery) (*models.PreviousCloseResponse, error)```


- Aggregates (Bar)
- - ```Aggregates(opts *models.AggregatesQuery) (*models.AggregatesResponse, error)```

#### Reference, Crypto, Forex will be added in the future.