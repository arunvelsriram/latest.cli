package common

import "net/http"

// HTTPClient interface that denotes the responsibilities of a HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
