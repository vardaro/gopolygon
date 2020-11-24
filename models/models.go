package models

// Bar corresponds to the "results" property in
// the response from /v2/aggs/
type Bar struct {
	O  float64 `json:"o"`
	H  float64 `json:"h"`
	L  float64 `json:"l"`
	C  float64 `json:"c"`
	V  float64 `json:"v"`
	VW float64 `json:"vw"`
	T  int64   `json:"t"`
	N  int64   `json:"n"`
}

// AggregatesResponse corresponds to the results from
// /v2/aggs
type AggregatesResponse struct {
	Symbol       string `json:"ticker"`
	Status       string `json:"status"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Bar  `json:"results"`
}
