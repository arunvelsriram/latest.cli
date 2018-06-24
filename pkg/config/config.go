package config

var version = "1.0.0-dev"

// GetVersion gets current app version
func GetVersion() string {
	return version
}

// SetVersion sets the current app version
func SetVersion(v string) {
	if v != "" {
		version = v
	}
}
