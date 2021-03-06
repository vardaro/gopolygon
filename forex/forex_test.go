package forex

// Although this is a test file, the test will often fail due
// being rate limited by the api when testing each function consecutively

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
		Symbol:     "C:EURUSD",
		Multiplier: 1,
		Timespan:   "day",
		From:       "2020-10-14",
		To:         "2020-10-14",
	}
	resp, err := client.Aggregates(query)
	if err != nil {
		t.Errorf("Error in client.Aggregate")
		return
	}
	expectedVol := 115644.00
	expectedResultsCount := 1
	if expectedResultsCount != resp.ResultsCount {
		t.Errorf("Aggregate.ResultsCount = %d; want %d", resp.ResultsCount, expectedResultsCount)
		return
	}

	if expectedVol != resp.Results[0].V {
		t.Errorf("Aggregate.Results[0].V = %v; want %v", resp.Results[0].V, expectedVol)
	}
}

func TestPreviousClose(t *testing.T) {
	unadjusted := true
	query := &models.PreviousCloseQuery{
		Symbol:     "C:EURUSD",
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

func TestGroupedDailyBar(t *testing.T) {
	query := &models.GroupedDailyBarsQuery{
		Date: "2020-10-14",
	}

	resp, err := client.GroupedDailyBars(query)
	if err != nil {
		t.Errorf("Error in client.GroupedDailyBars")
		return
	}

	expectedQueryCount := 1263
	if expectedQueryCount != resp.QueryCount {
		t.Errorf("PreviousClose.QueryCount = %v, want %v", resp.QueryCount, expectedQueryCount)
	}

}
