package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/vardaro/gopolygon/models"
)

// BaseURL of polygon
const BaseURL = "https://api.polygon.io"

// NewAPIError returns a Polygon API Error
func NewAPIError(response *http.Response) error {
	err := &models.APIError{}
	unmarshalPolygonResponse(response, &err)
	return err
}

// UnmarshalPolygonResponse Casts a Polygon response to interface
func unmarshalPolygonResponse(response *http.Response, data interface{}) error {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

// Get wrapper
// func Get(url *url.URL) (*http.Response, error) {
// 	fmt.Println(url)

// 	response, err := http.Get(url.String())

// 	// Catch API error
// 	if response.StatusCode != 200 {
// 		return nil, NewAPIError(response)
// 	}

// 	return response, err
// }

// Get wrapper
func Get(url *url.URL, responseStruct interface{}) (interface{}, error) {
	response, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	// Catch API error
	if response.StatusCode != 200 {
		return nil, NewAPIError(response)
	}

	err = unmarshalPolygonResponse(response, responseStruct)
	if err != nil {
		return nil, err
	}
	return responseStruct, nil
}
