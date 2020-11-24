package testing

import (
	"testing"

	"github.com/vardaro/gopolygon"
)

//TestClient test creating instance of polygon client.
func TestClient(t *testing.T) {
	want := "1001"
	got := gopolygon.NewClient("1001")

	if got.APIKey != want {
		t.Errorf("NewClient().APIKey = %v, want %v", got.APIKey, want)
	}
}
