package node

import (
	"fmt"

	"github.com/arunvelsriram/latest.cli/pkg/common"
)

// NPMRegistry exposes methods to talk to the npm registry
type NPMRegistry struct {
	url    string
	client common.JSONAPIClient
}

// LatestVersion gets the latest version of a node module
func (r *NPMRegistry) LatestVersion(name string) (string, error) {
	var version string

	uri := fmt.Sprintf("%s/%s/latest", r.url, name)
	data, err := r.client.GetJSON(uri)
	if err != nil {
		return version, err
	}

	if _, ok := data["version"]; !ok {
		return version, fmt.Errorf("Unable to get version from response data")
	}
	version = data["version"].(string)
	return version, nil
}

// NewNPMRegistry gives a new NPMRegistry
func NewNPMRegistry(url string, client common.JSONAPIClient) *NPMRegistry {
	return &NPMRegistry{url: url, client: client}
}
