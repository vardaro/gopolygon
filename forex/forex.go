package forex

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/vardaro/gopolygon/common"
	"github.com/vardaro/gopolygon/models"
)

const (
	routeAggregates       = "%v/v2/aggs/ticker/%v/range/%v/%v/%v/%v"
	routePreviousClose    = "%v/v2/aggs/ticker/%v/prev"
	routeGroupedDailyBars = "%v/v2/aggs/grouped/locale/global/market/fx/%v"
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
		fmt.Sprintf(routeAggregates, common.BaseURL, opts.Symbol, opts.Multiplier, opts.Timespan, opts.From, opts.To))
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
	response, err := common.Get(url)
	if err != nil {
		return nil, err
	}

	result := &models.AggregatesResponse{}
	err = common.UnmarshalPolygonResponse(response, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// PreviousClose function to query the Previous Close endpoint
func (c *Client) PreviousClose(opts *models.PreviousCloseQuery) (*models.PreviousCloseResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(routePreviousClose, common.BaseURL, opts.Symbol))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)

	if opts.Unadjusted != nil {
		q.Set("unadjusted", strconv.FormatBool(*opts.Unadjusted))
	}

	url.RawQuery = q.Encode()
	resp, err := common.Get(url)
	if err != nil {
		return nil, err
	}

	result := &models.PreviousCloseResponse{}
	err = common.UnmarshalPolygonResponse(resp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// GroupedDailyBars queries the GroupedDaily route
func (c *Client) GroupedDailyBars(opts *models.GroupedDailyBarsQuery) (*models.GroupedDailyBarsResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(routeGroupedDailyBars, common.BaseURL, opts.Date))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)

	if opts.Unadjusted != nil {
		q.Set("unadjusted", strconv.FormatBool(*opts.Unadjusted))
	}

	url.RawQuery = q.Encode()
	resp, err := common.Get(url)
	if err != nil {
		return nil, err
	}

	result := &models.GroupedDailyBarsResponse{}
	err = common.UnmarshalPolygonResponse(resp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
