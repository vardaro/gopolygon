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
	aggregatesURL = "%v/aggs/ticker/%v/range/%v/%v/%v/%v"
)

var (
	baseURL = "https://api.polygon.io/v2"
	get     = func(url *url.URL) (*http.Response, error) {
		return http.Get(url.String())
	}
)

// Aggregates corresponds to the /aggs/ route.
func (c *Client) Aggregates(stockTicker string, multiplier int, timespan string, from, to *time.Time, unadjusted *bool) (*AggregatesResponse, error) {
	// Build URL
	url, err := url.Parse(fmt.Sprintf(aggregatesURL, baseURL, stockTicker, multiplier, timespan, from.Unix()*1000, to.Unix()*1000))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)

	// Cast unadjusted bool -> string
	if unadjusted != nil {
		q.Set("unadjusted", strconv.FormatBool(*unadjusted))
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

// Casts a Polygon response to interface
func unmarshalPolygonResponse(response *http.Response, data interface{}) error {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}
