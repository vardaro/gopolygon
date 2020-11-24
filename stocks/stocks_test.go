package stocks

import (
	"fmt"
	"os"
	"testing"

	"github.com/vardaro/gopolygon/models"
)

var (
	apikey string
	client *Client
)

func init() {
	apikey := os.Getenv("POLYGON_API_KEY")
	if apikey == "" {
		fmt.Println("Cant find API Key")
		return
	}
	client = NewClient(apikey)
}

func TestClient(t *testing.T) {
	want := "1001"
	got := NewClient("1001")

	if got.APIKey != want {
		t.Errorf("NewClient().APIKey = %v, want %v", got.APIKey, want)
	}
}

// Test the stock aggregate function.
func TestAggregate(t *testing.T) {
	query := &models.AggregatesQuery{
		Symbol:     "AAPL",
		Multiplier: 1,
		Timespan:   "day",
		From:       "2020-01-01",
		To:         "2020-01-02",
	}
	resp, err := client.Aggregates(query)
	if err != nil {
		t.Errorf("Error in client.Aggregate")
		return
	}
	expectedVW := 74.6099
	expectedResultsCount := 1
	if expectedResultsCount != resp.ResultsCount {
		t.Errorf("Aggregate.ResultsCount = %d; want %d", resp.ResultsCount, expectedResultsCount)
		return
	}

	if expectedVW != resp.Results[0].VW {
		t.Errorf("Aggregate.Results[0].VW = %v; want %v", resp.Results[0].VW, expectedVW)
	}
}

func TestDailyOpenClose(t *testing.T) {
	query := &models.DailyOpenCloseQuery{
		Symbol: "AAPL",
		Date:   "2020-06-03",
	}
	resp, err := client.DailyOpenClose(query)
	if err != nil {
		t.Errorf("Error in client.DailyOpenClose")
		return
	}

	expectedHigh := 326.2

	if expectedHigh != resp.High {
		t.Errorf("DailyOpenClose.High = %v, want %v", resp.High, expectedHigh)
	}
}

func TestPreviousClose(t *testing.T) {
	unadjusted := true
	query := &models.PreviousCloseQuery{
		Symbol:     "AAPL",
		Unadjusted: &unadjusted,
	}
	resp, err := client.PreviousClose(query)
	if err != nil {
		t.Errorf("Error in client.PreviousClose")
		return
	}

	if unadjusted == resp.Adjusted {
		t.Errorf("PreviousClose.Adjusted = %v, want %v", resp.Adjusted, !resp.Adjusted)
	}
}
