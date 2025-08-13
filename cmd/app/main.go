package main

import (
	"github.com/xurvan/go-template/internal/config"
	"github.com/xurvan/go-template/internal/foo"
)

func main() {
	// Initialize configuration
	cfg := config.New()

	// Create and use the foo module
	fooModule := foo.New(cfg)
	fooModule.DoSomething()
}
