### Daily Open/Close
```golang
func (c *Client) PreviousClose(opts *models.PreviousCloseQuery) (*models.PreviousCloseResponse, error) 

type CryptoDailyOpenCloseQuery struct {
	From string
	To   string
	Date string
}

type CryptoDailyOpenCloseResponse struct {
	Symbol        string          `json:"symbol"`
	IsUTC         bool            `json:"isUTC"`
	Day           string          `json:"day"`
	Open          float64         `json:"open"`
	Close         float64         `json:"close"`
	OpenTrades    []HistoricTrade `json:"openTrades"`
	ClosingTrades []HistoricTrade `json:"ClosingTrades"`
}
```

### Grouped Daily (Bars)
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


### Previous Close
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

### Aggregates (Bar)
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