// Package foo provides example functionality with structured logging.
package foo

import (
	"context"
	"log/slog"
	"os"

	"github.com/tab-sama/go-template/internal/config"
)

// Foo performs example operations with structured logging.
type Foo struct {
	logger *slog.Logger
}

// New creates a new Foo instance with a configured stdout logger.
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

// DoSomething demonstrates logging at different levels.
func (f *Foo) DoSomething(ctx context.Context) {
	f.logger.DebugContext(ctx, "Debug message from foo")
	f.logger.InfoContext(ctx, "Info message from foo")
	f.logger.WarnContext(ctx, "Warning message from foo")
	f.logger.ErrorContext(ctx, "Error message from foo")
}
