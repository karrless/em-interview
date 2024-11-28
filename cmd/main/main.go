package main

import (
	"context"

	"github.com/karrless/em-interview/internal/config"
	"github.com/karrless/em-interview/pkg/db/postgres"
	"github.com/karrless/em-interview/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	mainLogger := logger.New(true)
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)
	mainLogger.Info("Starting application")

	cfg := config.New()
	if cfg == nil {
		mainLogger.Fatal("Failed to read config")
	}
	mainLogger.Debug("Config loaded", zap.Any("config", cfg))
	_, err := postgres.New(&ctx, cfg.PostgresConfig)
	if err != nil {
		mainLogger.Fatal("Failed to connect to database", zap.Error(err))
	}
	mainLogger.Debug("Database connected")
}
