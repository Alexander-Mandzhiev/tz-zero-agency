package handlers

import (
	"tz-zero-agency/internal/entity"
	category_repository "tz-zero-agency/internal/repository/category.repository"
	"tz-zero-agency/pkg/logger"
	"tz-zero-agency/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) createCategory(c *fiber.Ctx) error {
	token, err := h.token(c)
	if err != nil {
		return h.error(c, fiber.StatusUnauthorized, err)
	}
	var category *entity.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := validate.ValidateCategories(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.Categoryes.Create(h.ctx, category, token); err != nil {
		h.log.Error("Ошибка создания категории: %v", logger.Err(err))
		return h.error(c, fiber.StatusInternalServerError, category_repository.ErrCategoryExists)
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}
func (h *Handler) getCategories(c *fiber.Ctx) error {
	token, err := h.token(c)
	if err != nil {
		return h.error(c, fiber.StatusUnauthorized, err)
	}
	category, err := h.service.Categoryes.GetAll(h.ctx, token)
	if err != nil {
		h.log.Error("Ошибка при получении новостей: %v", logger.Err(err))
		return h.error(c, fiber.StatusInternalServerError, category_repository.ErrInvalidCredentials)
	}
	return h.respond(c, fiber.StatusOK, category)
}
