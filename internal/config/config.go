// Package config provides configuration management for the Go project template.
// It handles application configuration including log levels with environment variable support.
package config

import (
	"log/slog"
	"os"
	"strings"
)

// Config holds the application configuration settings.
// It contains settings for logging and other configurable aspects of the application.
type Config struct {
	LogLevel slog.Level
}

// New creates and returns a new Config instance with default settings.
// It sets the default log level to INFO and checks for the LOG_LEVEL environment variable
// to override the default. Supported values are DEBUG, INFO, WARN, and ERROR (case-insensitive).
func New() *Config {
	cfg := &Config{
		LogLevel: slog.LevelInfo, // Default log level
	}

	// Check for LOG_LEVEL environment variable
	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		switch strings.ToUpper(logLevel) {
		case "DEBUG":
			cfg.LogLevel = slog.LevelDebug
		case "INFO":
			cfg.LogLevel = slog.LevelInfo
		case "WARN":
			cfg.LogLevel = slog.LevelWarn
		case "ERROR":
			cfg.LogLevel = slog.LevelError
		}
	}

	return cfg
}
