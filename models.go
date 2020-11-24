package gopolygon

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
	Symbol       string `json:"ticker"`
	Status       string `json:"status"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Bar  `json:"results"`
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
	Symbol       string             `json:"ticker"`
	ResultsCount int                `json:"results_count"`
	Success      bool               `json:"success"`
	DBLatency    int                `json:"db_latency"`
	Results      []HistoricTrade    `json:"results"`
	Map          map[string]MapItem `json:"map"`
}

// APIError struct wrapper for API errors
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}
