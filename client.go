package fera

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// The current base URL of the free api.
const StandardBaseURL = "https://api.exchangeratesapi.io"

type Client struct {
	baseURL string
}

// Return a new Client object.
// This builder gives the possiblity to override the base URL if it will ever change.
func NewClient(base string) *Client {
	if base == "" {
		return &Client{
			baseURL: StandardBaseURL,
		}
	}

	return &Client{
		baseURL: base,
	}
}

func (c Client) getRequest(url string, ctx context.Context) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, reqErr := http.DefaultClient.Do(req)
	if reqErr != nil {
		return nil, reqErr
	}

	return res, nil
}

func (c Client) extractAndParseBody(res http.Response, parse interface{}) error {
	jsonBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if unerr := json.Unmarshal(jsonBody, parse); unerr != nil {
		return unerr
	}

	return nil
}
