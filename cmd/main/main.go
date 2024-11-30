package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/karrless/em-interview/internal/config"
	"github.com/karrless/em-interview/internal/repository"
	"github.com/karrless/em-interview/internal/service"
	"github.com/karrless/em-interview/internal/transport/rest"
	"github.com/karrless/em-interview/pkg/db/migrations"
	"github.com/karrless/em-interview/pkg/db/postgres"
	"github.com/karrless/em-interview/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	cfg := config.New("")
	if cfg == nil {
		log.Fatal("Failed to read config")
	}

	mainLogger := logger.New(cfg.Debug)
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)
	mainLogger.Info("Starting application")

	mainLogger.Debug("Config loaded", zap.Any("config", cfg))

	db, err := postgres.New(&ctx, cfg.PostgresConfig)
	if err != nil {
		mainLogger.Fatal("Failed to connect to database", zap.Error(err))
	}
	mainLogger.Debug("Database connected")

	migrationsVersion, err := migrations.Up(db.DB)
	if err != nil {
		mainLogger.Fatal("Failed to apply migrations", zap.Error(err))
	}
	if migrationsVersion == 0 {
		mainLogger.Debug("No new migrations")
	}
	mainLogger.Debug("Migrations applied", zap.Int("Migrate version", migrationsVersion))

	songsRepo := repository.NewSongRepository(db)
	externalAPIRepo := repository.NewExtarnalAPIRepository(&cfg.ExternalAPIConfig)

	songsService := service.NewSongsService(songsRepo, externalAPIRepo)

	server := rest.New(&ctx, cfg.ServerConfig, songsService, cfg.Debug)

	graceChannel := make(chan os.Signal, 1)
	signal.Notify(graceChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.Run(&ctx); err != nil {
			mainLogger.Fatal("failed to start server", zap.Error(err))
		}
	}()

	<-graceChannel
	db.Close()
	mainLogger.Debug("Database connection closed")
	mainLogger.Info("Graceful shutdown!")
}
