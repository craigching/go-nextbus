package services

import (
	"bytes"
	"net/http"
	"time"
)

// HTTPClient is a simple http client for use as a base REST client
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient makes a new HTTPClient
func NewHTTPClient() HTTPClient {
	return HTTPClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Get uses HTTP GET to get the contents of the url and returns them
// as a string
func (c *HTTPClient) Get(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/json")
	resp, err := c.client.Do(req)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	resp.Body.Close()

	return buf.String(), nil
}
