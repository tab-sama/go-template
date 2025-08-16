package main

import (
	"github.com/xurvan/go-template/internal/config"
	"github.com/xurvan/go-template/internal/foo"
)

func main() {
	cfg := config.New()

	f := foo.New(cfg)
	f.DoSomething()
}
