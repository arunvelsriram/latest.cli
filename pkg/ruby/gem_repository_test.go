package ruby_test

import (
	"fmt"
	"testing"

	"github.com/arunvelsriram/latest.cli/pkg/internal/mock"
	"github.com/arunvelsriram/latest.cli/pkg/ruby"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewGemRepository(t *testing.T) {
	repo := ruby.NewGemRepository("https://api-url", &mock.JSONAPIClient{})

	assert.NotNil(t, repo)
}

func TestLatestVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock.NewJSONAPIClient(ctrl)
	mockData := map[string]interface{}{"name": "rails", "version": "5.2.0"}
	mockClient.EXPECT().GetJSON("http://api-url/rails.json").Return(mockData, nil)
	repo := ruby.NewGemRepository("http://api-url", mockClient)

	actualVersion, err := repo.LatestVersion("rails")

	assert.Nil(t, err)
	assert.Equal(t, "5.2.0", actualVersion)
}

func TestLatestVersionShouldReturnErrorIfVersionNotFoundInJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock.NewJSONAPIClient(ctrl)
	mockData := map[string]interface{}{"name": "rails"}
	mockClient.EXPECT().GetJSON("http://api-url/rails.json").Return(mockData, nil)
	repo := ruby.NewGemRepository("http://api-url", mockClient)

	actualVersion, err := repo.LatestVersion("rails")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Unable to get version from response data")
	assert.Empty(t, actualVersion)
}

func TestShouldReturnErrorIfGetJSONThrowsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock.NewJSONAPIClient(ctrl)
	mockClient.EXPECT().GetJSON("http://api-url/rails.json").Return(nil, fmt.Errorf("Some error"))
	repo := ruby.NewGemRepository("http://api-url", mockClient)

	actualVersion, err := repo.LatestVersion("rails")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Some error")
	assert.Empty(t, actualVersion)
}
