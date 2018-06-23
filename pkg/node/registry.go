package node

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/arunvelsriram/latest.cli/pkg/common"
)

// Registry exposes methods to talk to the node registry
type Registry struct {
	url    string
	client common.HTTPClient
}

// LatestVersion gets the latest version of a node module
func (r *Registry) LatestVersion(nodeModule string) (string, error) {
	var data map[string]interface{}
	var latestVersion string

	uri := fmt.Sprintf("%s/%s/latest", r.url, nodeModule)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return latestVersion, err
	}

	res, err := r.client.Do(req)
	if err != nil {
		return latestVersion, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		return latestVersion, err
	}

	if _, ok := data["version"]; !ok {
		return latestVersion, errors.New("failed to get version from response")
	}
	latestVersion = data["version"].(string)
	return latestVersion, nil
}

// NewRegistry gives a new node registry
func NewRegistry(url string, client common.HTTPClient) *Registry {
	return &Registry{url: url, client: client}
}
