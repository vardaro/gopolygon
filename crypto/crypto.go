package crypto

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
	routeGroupedDailyBars = "%v/v2/aggs/grouped/locale/global/market/crypto/%v"
	routeDailyOpenClose   = "%v/v1/open-close/crypto/%v/%v/%v"
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

	inter, err := common.Get(url, &models.AggregatesResponse{})
	return inter.(*models.AggregatesResponse), err
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

	inter, err := common.Get(url, &models.PreviousCloseResponse{})
	return inter.(*models.PreviousCloseResponse), err

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

	inter, err := common.Get(url, &models.GroupedDailyBarsResponse{})
	return inter.(*models.GroupedDailyBarsResponse), err
}

// DailyOpenClose function to query the DailyOpenClose route
// https://polygon.io/docs/get_v1_open-close__stocksTicker___date__anchor
func (c *Client) DailyOpenClose(opts *models.CryptoDailyOpenCloseQuery) (*models.CryptoDailyOpenCloseResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(routeDailyOpenClose, common.BaseURL, opts.From, opts.To, opts.Date))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)
	url.RawQuery = q.Encode()

	inter, err := common.Get(url, &models.CryptoDailyOpenCloseResponse{})
	return inter.(*models.CryptoDailyOpenCloseResponse), err
}
