package config

import "fmt"

var majorVersion = "1.0"
var minorVersion = "dev"

// GetVersion gets current app version
func GetVersion() string {
	return fmt.Sprintf("%s.%s", majorVersion, minorVersion)
}

// SetVersion sets the current app version
func SetVersion(major string, minor string) {
	majorVersion = major
	minorVersion = minor
}
