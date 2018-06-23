package node_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/arunvelsriram/latest.cli/pkg/internal/mock"
	"github.com/arunvelsriram/latest.cli/pkg/node"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewNodeRegistry(t *testing.T) {
	registry := node.NewRegistry("https://registry.npmjs.org", &mock.HTTPClient{})

	assert.NotNil(t, registry)
}

func TestLatestVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		mockResponseData := ioutil.NopCloser(bytes.NewReader([]byte(`
			{"version": "6.1.0", "name": "npm", "description": "a package manager for JavaScript"}`,
		)))

		assert.Equal(t, "http://registry-base-url/npm/latest", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return &http.Response{Body: mockResponseData}, nil
	})
	registry := node.NewRegistry("http://registry-base-url", mockHTTPClient)

	actualVersion, err := registry.LatestVersion("npm")

	assert.Nil(t, err)
	assert.Equal(t, "6.1.0", actualVersion)
}

func TestLatestVersionWrongResponseData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		mockResponseData := ioutil.NopCloser(bytes.NewReader([]byte(`{"key": "value"}`)))

		assert.Equal(t, "http://registry-base-url/npm/latest", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return &http.Response{Body: mockResponseData}, nil
	})
	registry := node.NewRegistry("http://registry-base-url", mockHTTPClient)

	actualVersion, err := registry.LatestVersion("npm")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Unable to get version from response data")
	assert.Empty(t, actualVersion)
}

func TestLatestVersionErrorResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "http://registry-base-url/npm/latest", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return nil, errors.New("Some error")
	})
	registry := node.NewRegistry("http://registry-base-url", mockHTTPClient)

	actualVersion, err := registry.LatestVersion("npm")

	assert.NotNil(t, err)
	assert.Empty(t, actualVersion)
}
