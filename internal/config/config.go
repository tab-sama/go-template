package config

import (
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	LogLevel slog.Level
}

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
