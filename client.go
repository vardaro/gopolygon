package gopolygon

// EnvAPIKey Environment var that corresponds to api key
// Mostly here for testing purposes
const EnvAPIKey = "POLYGON_API_KEY"

// Client is polygon api client
type Client struct {
	APIKey string
}

// NewClient returns a new Client instance.
func NewClient(apikey string) *Client {
	return &Client{APIKey: apikey}
}
