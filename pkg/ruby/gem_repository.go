package ruby

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arunvelsriram/latest.cli/pkg/common"
)

// GemRepository exposes methods to talk to rubygems repository
type GemRepository struct {
	url    string
	client common.HTTPClient
}

// LatestVersion gets the latest version of a ruby gem
func (r *GemRepository) LatestVersion(name string) (string, error) {
	var data map[string]interface{}
	var latestVersion string

	uri := fmt.Sprintf("%s/%s.json", r.url, name)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return latestVersion, err
	}

	res, err := r.client.Do(req)
	if err != nil {
		return latestVersion, err
	}

	if res.StatusCode != http.StatusOK {
		return latestVersion, fmt.Errorf("Unable to fetch details for %s, StatusCode: %d", name, res.StatusCode)
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

// NewGemRepository gives a new GemRepository
func NewGemRepository(url string, client common.HTTPClient) *GemRepository {
	return &GemRepository{url: url, client: client}
}
