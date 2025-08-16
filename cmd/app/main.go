// Package main is the entry point for the application.
package main

import (
	"context"

	"github.com/xurvan/go-template/internal/config"
	"github.com/xurvan/go-template/internal/foo"
)

func main() {
	cfg := config.New()

	ctx := context.Background()

	f := foo.New(cfg)
	f.DoSomething(ctx)
}
