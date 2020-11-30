# gopolygon
Go SDK for Polygon API.

## Installation
```go get github.com/vardaro/gopolygon```

## Example

### REST Example

```golang
package main

import (
	"fmt"

	"github.com/vardaro/gopolygon/models"
	"github.com/vardaro/gopolygon/stocks"
)

// Print the closing price of AAPL on 3/20/2020
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

## Documentation

### Stocks

- [Stocks SDK Documentation](./stocks/docs.md)

### Forex

- [Forex SDK Documenation](./forex/docs.md)

#### Reference, Crypto, Forex will be added in the future. 