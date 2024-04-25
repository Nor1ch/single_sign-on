package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"sso/pkg/logger/slogpretty"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	cfg.Log()

	logger := slogpretty.New(cfg.Env)
	logger.Info("service starting")

	application := app.New(logger, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	go application.GRPCServer.MustRun()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	logger.Info("stopping application", slog.String("signal", sign.String()))

	application.GRPCServer.Stop()
	logger.Info("application stopped")
}
