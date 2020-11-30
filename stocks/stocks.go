package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/vardaro/gopolygon/models"
)

const (
	baseURL = "https://api.polygon.io"

	routeAggregates       = "%v/v2/aggs/ticker/%v/range/%v/%v/%v/%v"
	routeHistoricTrades   = "%v/v2/ticks/stocks/trades/%v/%v"
	routeDailyOpenClose   = "%v/v1/open-close/%v/%v"
	routePreviousClose    = "%v/v2/aggs/ticker/%v/prev"
	routeGroupedDailyBars = "%v/v2/aggs/grouped/locale/us/market/stocks/%v"
)

// Client is polygon api client
type Client struct {
	APIKey string
}

// NewClient returns a new Client instance.
func NewClient(apikey string) *Client {
	return &Client{APIKey: apikey}
}

// Aggregates corresponds to the /aggs/ route.
// https://polygon.io/docs/get_v2_aggs_ticker__stocksTicker__range__multiplier___timespan___from___to__anchor
func (c *Client) Aggregates(opts *models.AggregatesQuery) (*models.AggregatesResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(routeAggregates, baseURL, opts.Symbol, opts.Multiplier, opts.Timespan, opts.From, opts.To))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)

	// Cast unadjusted bool -> string
	if opts.Unadjusted != nil {
		q.Set("unadjusted", strconv.FormatBool(*opts.Unadjusted))
	}

	url.RawQuery = q.Encode()
	response, err := get(url)
	if err != nil {
		return nil, err
	}
	result := &models.AggregatesResponse{}
	err = unmarshalPolygonResponse(response, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HistoricTrades queries the Historic Trades route.
// Don't have a paid subscription to adequately test this
// https://polygon.io/docs/get_v2_ticks_stocks_trades__ticker___date__anchor
func (c *Client) HistoricTrades(opts *models.HistoricTradesQuery) (*models.HistoricTradesResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(routeHistoricTrades, baseURL, opts.Symbol, opts.Date))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)
	url.RawQuery = q.Encode()

	// Set other params if applicable
	if opts.Timestamp != nil {
		q.Set("timestamp", strconv.FormatInt(*opts.Timestamp, 10))
	}

	if opts.TimestampLimit != nil {
		q.Set("timestampLimit", strconv.FormatInt(*opts.TimestampLimit, 10))
	}

	if opts.Reverse != nil {
		q.Set("reverse", strconv.FormatBool(*opts.Reverse))
	}

	if opts.Limit != nil {
		q.Set("limit", strconv.FormatInt(*opts.Limit, 10))
	}

	resp, err := get(url)
	if err != nil {
		return nil, err
	}

	result := &models.HistoricTradesResponse{}
	err = unmarshalPolygonResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// DailyOpenClose function to query the DailyOpenClose route
// https://polygon.io/docs/get_v1_open-close__stocksTicker___date__anchor
func (c *Client) DailyOpenClose(opts *models.DailyOpenCloseQuery) (*models.DailyOpenCloseResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(routeDailyOpenClose, baseURL, opts.Symbol, opts.Date))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)
	url.RawQuery = q.Encode()
	resp, err := get(url)

	if err != nil {
		return nil, err
	}

	result := &models.DailyOpenCloseResponse{}
	err = unmarshalPolygonResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// PreviousClose function to query the Previous Close endpoint
func (c *Client) PreviousClose(opts *models.PreviousCloseQuery) (*models.PreviousCloseResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(routePreviousClose, baseURL, opts.Symbol))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)

	if opts.Unadjusted != nil {
		q.Set("unadjusted", strconv.FormatBool(*opts.Unadjusted))
	}

	url.RawQuery = q.Encode()
	resp, err := get(url)
	if err != nil {
		return nil, err
	}

	result := &models.PreviousCloseResponse{}
	err = unmarshalPolygonResponse(resp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// GroupedDailyBars queries the GroupedDaily route
func (c *Client) GroupedDailyBars(opts *models.GroupedDailyBarsQuery) (*models.GroupedDailyBarsResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(routeGroupedDailyBars, baseURL, opts.Date))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)

	if opts.Unadjusted != nil {
		q.Set("unadjusted", strconv.FormatBool(*opts.Unadjusted))
	}

	url.RawQuery = q.Encode()
	resp, err := get(url)
	if err != nil {
		return nil, err
	}

	result := &models.GroupedDailyBarsResponse{}
	err = unmarshalPolygonResponse(resp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Casts a Polygon response to interface
func unmarshalPolygonResponse(response *http.Response, data interface{}) error {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

func get(url *url.URL) (*http.Response, error) {
	fmt.Println(url)
	return http.Get(url.String())
}
