package config

import "fmt"

var majorVersion = 0
var minorVersion = 0
var patchVersion = 1

// GetVersion gets current app version
func GetVersion() string {
	return fmt.Sprintf("%d.%d.%d", majorVersion, minorVersion, patchVersion)
}

// SetVersion sets the current app version
func SetVersion(major int, minor int, patch int) {
	majorVersion = major
	minorVersion = minor
	patchVersion = patch
}
