package ruby_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/arunvelsriram/latest.cli/pkg/internal/mock"
	"github.com/arunvelsriram/latest.cli/pkg/ruby"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewGemRepository(t *testing.T) {
	repo := ruby.NewGemRepository("https://repo-api-url", &mock.HTTPClient{})

	assert.NotNil(t, repo)
}

func TestGemLatestVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		mockResponseData := ioutil.NopCloser(bytes.NewReader([]byte(`
			{"version": "5.2.0", "name": "rails"}`,
		)))

		assert.Equal(t, "http://repo-api-url/rails.json", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return &http.Response{Body: mockResponseData, StatusCode: http.StatusOK}, nil
	})
	registry := ruby.NewGemRepository("http://repo-api-url", mockHTTPClient)

	actualVersion, err := registry.LatestVersion("rails")

	assert.Nil(t, err)
	assert.Equal(t, "5.2.0", actualVersion)
}

func TestGemLatestVersionWrongResponseData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		mockResponseData := ioutil.NopCloser(bytes.NewReader([]byte(`{"key": "value"}`)))

		assert.Equal(t, "http://repo-api-url/rails.json", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return &http.Response{Body: mockResponseData, StatusCode: http.StatusOK}, nil
	})
	repo := ruby.NewGemRepository("http://repo-api-url", mockHTTPClient)

	actualVersion, err := repo.LatestVersion("rails")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Unable to get version from response data")
	assert.Empty(t, actualVersion)
}

func TestGemLatestVersionErrorMakingHTTPRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "http://repo-api-url/rails.json", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return nil, errors.New("Some error")
	})
	repo := ruby.NewGemRepository("http://repo-api-url", mockHTTPClient)

	actualVersion, err := repo.LatestVersion("rails")

	assert.NotNil(t, err)
	assert.Empty(t, actualVersion)
}
func TestGemLatestVersionResponseResponseStatusNotOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "http://repo-api-url/rails.json", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return &http.Response{StatusCode: http.StatusInternalServerError}, nil
	})
	registry := ruby.NewGemRepository("http://repo-api-url", mockHTTPClient)

	actualVersion, err := registry.LatestVersion("rails")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Unable to fetch details for rails, StatusCode: 500")
	assert.Empty(t, actualVersion)
}
