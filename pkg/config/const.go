package config

import "os"

// Build-time variables (set via ldflags)
var (
	Version   = "dev"     // Application version
	GitCommit = "unknown" // Git commit hash
	BuildDate = "unknown" // Build timestamp
)

// Environment variables
var (
	// APP_ENV: Application environment (dev, staging, prod)
	AppEnv = getEnv("APP_ENV", "dev")

	// APP_VERSION: Override version from environment (optional)
	AppVersion = getEnv("APP_VERSION", Version)
)

// Constants
const (
	LogLevelDebug   string = "debug"
	LogLevelInfo    string = "info"
	LogLevelWarning string = "warn"
	LogLevelError   string = "error"

	LogFilePath string = "log_output.txt"
)

// getEnv retrieves environment variable with fallback default
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
