// Package foo provides example functionality for the Go project template.
// It demonstrates structured logging using slog with configurable log levels.
package foo

import (
	"context"
	"log/slog"
	"os"

	"github.com/xurvan/go-template/internal/config"
)

// Foo represents a component that performs example operations with structured logging.
// It encapsulates a logger instance configured according to the provided configuration.
type Foo struct {
	logger *slog.Logger
}

// New creates and returns a new Foo instance with a configured logger.
// The logger is set up with a text handler that writes to stdout,
// using the log level specified in the provided configuration.
func New(cfg *config.Config) *Foo {
	logHandler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level:       cfg.LogLevel,
			AddSource:   true,
			ReplaceAttr: nil,
		},
	)

	logger := slog.New(logHandler)

	return &Foo{
		logger: logger,
	}
}

// DoSomething demonstrates logging at different levels (Debug, Info, Warn, Error).
// It takes a context and logs example messages at each log level to showcase
// the structured logging capabilities of the component.
func (f *Foo) DoSomething(ctx context.Context) {
	f.logger.DebugContext(ctx, "Debug message from foo")
	f.logger.InfoContext(ctx, "Info message from foo")
	f.logger.WarnContext(ctx, "Warning message from foo")
	f.logger.ErrorContext(ctx, "Error message from foo")
}
