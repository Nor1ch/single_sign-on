package main

import (
	"sso/internal/config"
	"sso/pkg/logger/slogpretty"
)

func main() {
	cfg := config.MustLoad()
	cfg.Log()

	logger := slogpretty.New(cfg.Env)
	logger.Info("service starting")
}
