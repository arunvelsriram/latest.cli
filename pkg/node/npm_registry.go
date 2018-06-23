package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arunvelsriram/latest.cli/pkg/common"
)

// NPMRegistry exposes methods to talk to the npm registry
type NPMRegistry struct {
	url    string
	client common.HTTPClient
}

// LatestVersion gets the latest version of a node module
func (r *NPMRegistry) LatestVersion(nodeModule string) (string, error) {
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

	if res.StatusCode != http.StatusOK {
		return latestVersion, fmt.Errorf("Unable to fetch details for %s, StatusCode: %d", nodeModule, res.StatusCode)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		return latestVersion, err
	}

	if _, ok := data["version"]; !ok {
		return latestVersion, fmt.Errorf("Unable to get version from response data")
	}
	latestVersion = data["version"].(string)
	return latestVersion, nil
}

// NewNPMRegistry gives a new NPMRegistry
func NewNPMRegistry(url string, client common.HTTPClient) *NPMRegistry {
	return &NPMRegistry{url: url, client: client}
}
