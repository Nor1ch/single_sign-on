package main

import (
	"sso/internal/app"
	"sso/internal/config"
	"sso/pkg/logger/slogpretty"
)

func main() {
	cfg := config.MustLoad()
	cfg.Log()

	logger := slogpretty.New(cfg.Env)
	logger.Info("service starting")

	application := app.New(logger, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	application.GRPCServer.MustRun()
}
