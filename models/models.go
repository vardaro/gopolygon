package models

// Bar corresponds to the "results" property in
// the response from /v2/aggs/
type Bar struct {
	O      float64 `json:"o"`
	H      float64 `json:"h"`
	L      float64 `json:"l"`
	C      float64 `json:"c"`
	V      float64 `json:"v"`
	VW     float64 `json:"vw"`
	T      int64   `json:"t"`
	N      int64   `json:"n"`
	Ticker string  `json:"T"`
}

// HistoricTrade contains json of a trade from polygon
// The object contains both and 'I' and an 'i' field, so HistoricalTrade.OriginalID maps to 'I'
// and HistoricTrade.I maps to 'i'
type HistoricTrade struct {
	OriginalID int     `json:"ID"`
	T          int     `json:"t"`
	Y          int     `json:"y"`
	F          int     `json:"f"`
	Q          int     `json:"q"`
	I          string  `json:"i"` // Trade ID
	X          int     `json:"x"`
	S          int     `json:"s"`
	C          []int   `json:"c"`
	P          float64 `json:"p"` // float64, as it's a trade price
	Z          int     `json:"z"`
}

// AggregatesResponse corresponds to the results from
// /v2/aggs
type AggregatesResponse struct {
	Ticker       string `json:"ticker"`
	Status       string `json:"status"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Bar  `json:"results"`
}

// AggregatesQuery input payload for making queries to aggregates endpoint
// pointer fields are optional, so they can be nil
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

// MapItem corresponds to an element within the map property of a HistoricTrades Response
// Maps can be found here, https://polygon.io/docs/get_v2_ticks_stocks_trades__ticker___date__anchor
type MapItem struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// HistoricTradesResponse corresponds to results from
// /ticks/
type HistoricTradesResponse struct {
	Ticker       string             `json:"ticker"`
	ResultsCount int                `json:"results_count"`
	Success      bool               `json:"success"`
	DBLatency    int                `json:"db_latency"`
	Results      []HistoricTrade    `json:"results"`
	Map          map[string]MapItem `json:"map"`
}

// HistoricTradesQuery query object for the HistoricTrades endpoint
type HistoricTradesQuery struct {
	Symbol         string
	Date           string
	Timestamp      *int64
	TimestampLimit *int64
	Reverse        *bool
	Limit          *int64
}

// DailyOpenCloseResponse response object of the Daily Open/Close route
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

// DailyOpenCloseQuery struct for querying the route
// why is polygon naming so inconsistent?
type DailyOpenCloseQuery struct {
	Symbol string
	Date   string // I think this doesnt support milliseconds?
}

// PreviousCloseResponse struct for PreviousClose endpoint
type PreviousCloseResponse struct {
	Ticker       string `json:"ticker"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Bar  `json:"results"`
}

// PreviousCloseQuery struct for querying the Previous Close route.
type PreviousCloseQuery struct {
	Symbol     string
	Unadjusted *bool
}

// GroupedDailyBarsResponse struct for grouped daily route resp
type GroupedDailyBarsResponse struct {
	Status       string `json:"status"`
	Adjusted     bool   `json:"adjusted"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	RequestID    string `json:"request_id"`
	Results      []Bar  `json:"results"`
}

// GroupedDailyBarsQuery struct for querying grouped daily route.
type GroupedDailyBarsQuery struct {
	Date       string
	Unadjusted *bool
}

// APIError struct wrapper for API errors
type APIError struct {
	Status  string `json:"status"`
	Message string `json:"error"`
}

func (e *APIError) Error() string {
	return e.Message
}
