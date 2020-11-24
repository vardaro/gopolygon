package gopolygon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
func (c *Client) Aggregates(stockTicker string, multiplier int, timespan string, from, to *time.Time, unadjusted bool) (*AggregatesResponse, error) {
	// Build URL
	url, err := url.Parse(fmt.Sprintf(aggregatesURL, baseURL, stockTicker, multiplier, timespan, from.Unix()*1000, to.Unix()*1000))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)

	// Cast unadjusted bool -> string
	if unadjusted {
		q.Set("unadjusted", strconv.FormatBool(unadjusted))
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
func (c *Client) HistoricTrades(stockTicker string, date *time.Time, timestamp int64, timestampLimit int64, reverse bool, limit int64) (*HistoricTradesResponse, error) {
	// Build URL
	url, err := url.Parse(fmt.Sprintf(historicTradesURL, baseURL, stockTicker, date.Unix()*1000))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)
	url.RawQuery = q.Encode()

	// Set other params if applicable
	if timestamp != 0 {
		q.Set("timestamp", strconv.FormatInt(timestamp, 10))
	}

	if timestampLimit != 0 {
		q.Set("timestampLimit", strconv.FormatInt(timestampLimit, 10))
	}

	if reverse {
		q.Set("reverse", strconv.FormatBool(reverse))
	}

	if limit != 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
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
