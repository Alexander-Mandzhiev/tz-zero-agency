package main

import (
	"context"
	"log/slog"
	"tz-zero-agency/internal/apiserver"
	"tz-zero-agency/internal/config"
	"tz-zero-agency/internal/repository"
	"tz-zero-agency/internal/service"
	"tz-zero-agency/pkg/logger"
	"tz-zero-agency/pkg/postgres"
)

func main() {
	log := logger.SetupLogger()
	log.Debug("Debug messages are enabled")

	cfg := config.NewConfig()
	log.Info("Initialize config", slog.String("port", cfg.Address))

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Error("failed to initialize db: %s", logger.Err(err))
	}
	defer postgres.CloseDB()
	log.Debug("Initialize database")

	repo := repository.NewRepository(log, db)
	log.Debug("Initialize repository")

	service := service.NewService(repo, log, cfg.TokenTTL, cfg.SecretKey)
	log.Debug("Initialize service")

	ctx := context.Background()

	app := apiserver.Init(ctx, log, service)
	log.Info("Starting app", slog.String("address", cfg.Address))

	if err := app.Listen(cfg.Address); err != nil {
		log.Error("error occured while running http server: %s", logger.Err(err))
	}
}
