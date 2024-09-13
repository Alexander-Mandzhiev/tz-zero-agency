package apiserver

import (
	"context"
	"log/slog"
	"tz-zero-agency/internal/handlers"
	"tz-zero-agency/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Init(ctx context.Context, log *slog.Logger, service *service.Service) *fiber.App {
	app := fiber.New()
	initRouters(ctx, log, app, service)
	return app
}
func initRouters(ctx context.Context, log *slog.Logger, app *fiber.App, service *service.Service) {
	handler := handlers.NewHandler(service, app, log, ctx)
	handler.InitRouters()
}
