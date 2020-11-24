package testing

import (
	"os"
	"testing"
	"time"

	"github.com/vardaro/gopolygon"
)

// Test the stock aggregate function.
func TestAggregate(t *testing.T) {
	apikey := os.Getenv(gopolygon.EnvAPIKey)
	if apikey == "" {
		t.Log("Cant find API Key")
		return
	}
	client := gopolygon.NewClient(apikey)

	jan1, _ := time.Parse("2006-01-02 15:04", "2020-01-01 00:00")
	jan2, _ := time.Parse("2006-01-02 15:04", "2020-01-02 00:00")

	resp, err := client.Aggregates("AAPL", 1, "day", &jan1, &jan2, nil)
	if err != nil {
		t.Errorf("Error in client.Aggregate")
	}
	expectedVW := 73.0982
	expectedResultsCount := 1
	if expectedResultsCount != resp.ResultsCount {
		t.Errorf("Aggregate.ResultsCount = %d; want %d", resp.ResultsCount, expectedResultsCount)
	}

	if expectedVW != resp.Results[0].VW {
		t.Errorf("Aggregate.Results[0].VW = %v; want %v", resp.Results[0].VW, expectedVW)
	}
}
