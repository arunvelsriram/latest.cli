package node_test

import (
	"fmt"
	"testing"

	"github.com/arunvelsriram/latest.cli/pkg/internal/mock"
	"github.com/arunvelsriram/latest.cli/pkg/node"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewNPMRegistry(t *testing.T) {
	registry := node.NewNPMRegistry("https://registry-base-url", &mock.JSONAPIClient{})

	assert.NotNil(t, registry)
}

func TestLatestVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock.NewJSONAPIClient(ctrl)
	mockData := map[string]interface{}{"name": "npm", "version": "6.1.0"}
	mockClient.EXPECT().GetJSON("http://api-url/npm/latest").Return(mockData, nil)
	registry := node.NewNPMRegistry("http://api-url", mockClient)

	actualVersion, err := registry.LatestVersion("npm")

	assert.Nil(t, err)
	assert.Equal(t, "6.1.0", actualVersion)
}

func TestLatestShouldReturnErronIfVersionNotFoundInJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock.NewJSONAPIClient(ctrl)
	mockData := map[string]interface{}{"name": "npm"}
	mockClient.EXPECT().GetJSON("http://api-url/npm/latest").Return(mockData, nil)
	registry := node.NewNPMRegistry("http://api-url", mockClient)

	actualVersion, err := registry.LatestVersion("npm")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Unable to get version from response data")
	assert.Empty(t, actualVersion)
}

func TestShouldReturnErrorIfGetJSONThrowsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock.NewJSONAPIClient(ctrl)
	mockClient.EXPECT().GetJSON("http://api-url/npm/latest").Return(nil, fmt.Errorf("Some error"))
	registry := node.NewNPMRegistry("http://api-url", mockClient)

	actualVersion, err := registry.LatestVersion("npm")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Some error")
	assert.Empty(t, actualVersion)
}
