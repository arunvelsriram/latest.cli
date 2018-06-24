package common_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/arunvelsriram/latest.cli/pkg/common"
	"github.com/arunvelsriram/latest.cli/pkg/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewPkgRegistryClient(t *testing.T) {
	repo := common.NewPkgRegistryClient(&mock.HTTPClient{})

	assert.NotNil(t, repo)
}

func TestGetJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		mockResponseData := ioutil.NopCloser(bytes.NewReader([]byte(`{"name": "rails"}`)))

		assert.Equal(t, "https://api-url", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return &http.Response{Body: mockResponseData, StatusCode: http.StatusOK}, nil
	})
	expectedData := map[string]interface{}{"name": "rails"}
	repo := common.NewPkgRegistryClient(mockHTTPClient)

	actualData, err := repo.GetJSON("https://api-url")

	assert.Nil(t, err)
	assert.Equal(t, expectedData, actualData)
}

func TestGetJSONShouldReturnErrorWhenHTTPRequestFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://api-url", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return nil, fmt.Errorf("Some error")
	})
	repo := common.NewPkgRegistryClient(mockHTTPClient)

	data, err := repo.GetJSON("https://api-url")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Some error")
	assert.Nil(t, data)
}

func TestGetJSONShouldReturnErrorWhenHTTPResponseStatusIsNotOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "https://api-url", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return &http.Response{StatusCode: http.StatusInternalServerError}, nil
	})
	repo := common.NewPkgRegistryClient(mockHTTPClient)

	data, err := repo.GetJSON("https://api-url")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Unable to fetch details, StatusCode: 500")
	assert.Nil(t, data)
}

func TestGetJSONShouldReturnErrorForInvalidJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHTTPClient := mock.NewHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		mockResponseData := ioutil.NopCloser(bytes.NewReader([]byte("not a JSON response")))

		assert.Equal(t, "https://api-url", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)

		return &http.Response{StatusCode: http.StatusOK, Body: mockResponseData}, nil
	})
	repo := common.NewPkgRegistryClient(mockHTTPClient)

	data, err := repo.GetJSON("https://api-url")

	assert.NotNil(t, err)
	assert.Nil(t, data)
}
