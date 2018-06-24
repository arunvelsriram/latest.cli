package config_test

import (
	"testing"

	"github.com/arunvelsriram/latest.cli/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestGetDefaultVersion(t *testing.T) {
	expectedVersion := "1.0.0-dev"

	actualVersion := config.GetVersion()

	assert.Equal(t, expectedVersion, actualVersion)
}

func TestSetVersionShouldIgnoreIfVersionIsEmpty(t *testing.T) {
	config.SetVersion("")
	expectedVersion := "1.0.0-dev"

	actualVersion := config.GetVersion()

	assert.Equal(t, expectedVersion, actualVersion)
}

func TestGetVersion(t *testing.T) {
	config.SetVersion("1.2.3")
	expectedVersion := "1.2.3"

	actualVersion := config.GetVersion()

	assert.Equal(t, expectedVersion, actualVersion)
}
