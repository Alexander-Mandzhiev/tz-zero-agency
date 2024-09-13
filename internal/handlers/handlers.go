package handlers

import (
	"context"
	"log/slog"
	"tz-zero-agency/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *service.Service
	app     *fiber.App
	log     *slog.Logger
	ctx     context.Context
}

func NewHandler(service *service.Service, app *fiber.App, log *slog.Logger, ctx context.Context) *Handler {
	return &Handler{service: service, app: app, log: log, ctx: ctx}
}

func (h *Handler) InitRouters() {

	h.app.Post("/users/signin", h.signin)
	h.app.Post("/users/signup", h.signup)

	h.app.Get("/news", h.getNews)
	h.app.Post("/news", h.createNews)
	h.app.Patch("/news/:id", h.UpdateNews)

	h.app.Get("/categories", h.getCategories)
	h.app.Post("/categories", h.createCategory)
}
