package foo

import (
	"log/slog"
	"os"

	"github.com/xurvan/go-template/internal/config"
)

type Foo struct {
	logger *slog.Logger
}

func New(cfg *config.Config) *Foo {
	// Create a logger with the configured log level
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.LogLevel,
	})

	logger := slog.New(logHandler)

	return &Foo{
		logger: logger,
	}
}

func (f *Foo) DoSomething() {
	f.logger.Debug("Debug message from foo module")
	f.logger.Info("Info message from foo module")
	f.logger.Warn("Warning message from foo module")
	f.logger.Error("Error message from foo module")
}
