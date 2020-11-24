package gopolygon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	aggregatesURL     = "%v/aggs/ticker/%v/range/%v/%v/%v/%v"
	historicTradesURL = "%v/ticks/stocks/trades/%v/%v"
)

var (
	baseURL = "https://api.polygon.io/v2"
	get     = func(url *url.URL) (*http.Response, error) {
		return http.Get(url.String())
	}
)

// Aggregates corresponds to the /aggs/ route.
func (c *Client) Aggregates(opts *AggregatesQuery) (*AggregatesResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(aggregatesURL, baseURL, opts.Symbol, opts.Multiplier, opts.Timespan, opts.From.Unix()*1000, opts.To.Unix()*1000))
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

	if response.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("error %v", response.StatusCode)
	}

	result := &AggregatesResponse{}
	err = unmarshalPolygonResponse(response, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HistoricTrades queries the Historic Trades route.
func (c *Client) HistoricTrades(opts *HistoricTradesQuery) (*HistoricTradesResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(historicTradesURL, baseURL, opts.Symbol, opts.Date.Unix()*1000))
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

	if resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("error %v", resp.StatusCode)
	}

	result := &HistoricTradesResponse{}
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
