package ruby

import (
	"fmt"

	"github.com/arunvelsriram/latest.cli/pkg/common"
)

// GemRepository exposes methods to talk to rubygems repository
type GemRepository struct {
	url    string
	client common.JSONAPIClient
}

// LatestVersion gets the latest version of a ruby gem
func (r *GemRepository) LatestVersion(name string) (string, error) {
	var version string

	uri := fmt.Sprintf("%s/%s.json", r.url, name)
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

// NewGemRepository gives a new GemRepository
func NewGemRepository(url string, client common.JSONAPIClient) *GemRepository {
	return &GemRepository{url: url, client: client}
}
