package handlers

import (
	"errors"
	"log"
	"strconv"
	"tz-zero-agency/internal/entity"
	news_repository "tz-zero-agency/internal/repository/news.repository"
	"tz-zero-agency/pkg/logger"
	"tz-zero-agency/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) getNews(c *fiber.Ctx) error {
	token, err := h.token(c)
	if err != nil {
		return h.error(c, fiber.StatusUnauthorized, err)
	}

	limitStr := c.Query("limit", "10")
	offsetStr := c.Query("offset", "0")

	news, err := h.service.News.GetAll(h.ctx, token, limitStr, offsetStr)
	if err != nil {
		h.log.Error("Ошибка при получении новостей: %v", logger.Err(err))
		return h.error(c, fiber.StatusInternalServerError, news_repository.ErrInvalidCredentials)
	}
	return h.respond(c, fiber.StatusOK, news)
}

func (h *Handler) createNews(c *fiber.Ctx) error {
	token, err := h.token(c)
	if err != nil {
		return h.error(c, fiber.StatusUnauthorized, err)
	}

	var news *entity.News
	if err := c.BodyParser(&news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if err := validate.ValidateNews(news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.News.Create(h.ctx, news, token); err != nil {
		h.log.Error("Ошибка создания новости: %v", logger.Err(err))
		return h.error(c, fiber.StatusInternalServerError, news_repository.ErrCreateNews)
	}

	return c.Status(fiber.StatusCreated).JSON(news)
}

func (h *Handler) UpdateNews(c *fiber.Ctx) error {
	op := "handlers.UpdateNews"
	token, err := h.token(c)
	if err != nil {
		return h.error(c, fiber.StatusUnauthorized, err)
	}
	idx := c.Params("id")

	var news *entity.News
	if err := c.BodyParser(&news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	id, err := strconv.Atoi(idx)
	if err != nil {
		log.Printf("%s: неверное значение для идентификатора: %v", op, err)
		return h.error(c, fiber.StatusBadRequest, errors.New("указано неверное значение для идентификатора, пожалуйста, введите корректное целое число"))
	}

	n, err := h.service.News.Update(h.ctx, token, news, id)
	if err != nil {
		h.log.Error("Ошибка создания новости: %v", logger.Err(err))
		return h.error(c, fiber.StatusInternalServerError, news_repository.ErrCreateNews)
	}

	return c.Status(fiber.StatusCreated).JSON(n)
}
