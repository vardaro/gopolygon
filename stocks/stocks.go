package stocks

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/vardaro/gopolygon/common"
	"github.com/vardaro/gopolygon/models"
)

const (
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

// HistoricTrades queries the Historic Trades route.
// Don't have a paid subscription to adequately test this
// https://polygon.io/docs/get_v2_ticks_stocks_trades__ticker___date__anchor
func (c *Client) HistoricTrades(opts *models.HistoricTradesQuery) (*models.HistoricTradesResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(routeHistoricTrades, common.BaseURL, opts.Symbol, opts.Date))
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

	inter, err := common.Get(url, &models.HistoricTradesResponse{})
	return inter.(*models.HistoricTradesResponse), err
}

// DailyOpenClose function to query the DailyOpenClose route
// https://polygon.io/docs/get_v1_open-close__stocksTicker___date__anchor
func (c *Client) DailyOpenClose(opts *models.DailyOpenCloseQuery) (*models.DailyOpenCloseResponse, error) {
	// Build URL
	url, err := url.Parse(
		fmt.Sprintf(routeDailyOpenClose, common.BaseURL, opts.Symbol, opts.Date))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("apiKey", c.APIKey)
	url.RawQuery = q.Encode()

	inter, err := common.Get(url, &models.DailyOpenCloseResponse{})
	return inter.(*models.DailyOpenCloseResponse), err
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
