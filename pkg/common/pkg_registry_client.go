package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// PkgRegistryClient contains method to fetch data from any package registry
type PkgRegistryClient struct {
	client HTTPClient
}

// GetJSON fetches JSON data from the given URL
func (c *PkgRegistryClient) GetJSON(url string) (map[string]interface{}, error) {
	var data map[string]interface{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return data, err
	}

	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("Unable to fetch details, StatusCode: %d", res.StatusCode)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		return data, err
	}

	return data, nil
}

// NewPkgRegistryClient returns a new PkgRegistryClient
func NewPkgRegistryClient(client HTTPClient) *PkgRegistryClient {
	return &PkgRegistryClient{client: client}
}
