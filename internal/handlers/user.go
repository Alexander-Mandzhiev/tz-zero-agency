package handlers

import (
	"log"
	"tz-zero-agency/internal/entity"
	user_repository "tz-zero-agency/internal/repository/user.repository"
	"tz-zero-agency/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) signin(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var req *entity.UserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}
	log.Printf("Received: %+v\n", req)

	token, err := h.service.User.Signin(h.ctx, req.Email, req.Password)
	if err != nil {
		h.log.Error("Ошибка при входе пользователя: %v", logger.Err(err))
		return h.error(c, fiber.StatusInternalServerError, user_repository.ErrUserExists)
	}
	return h.respond(c, fiber.StatusOK, map[string]string{"token": token})

}

func (h *Handler) signup(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var req *entity.UserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	id, err := h.service.User.Signup(h.ctx, req.Email, req.Password)
	if err != nil {
		h.log.Error("Ошибка при регистрации пользователя: %v", logger.Err(err))
		return h.error(c, fiber.StatusInternalServerError, user_repository.ErrUserExists)
	}

	return h.respond(c, fiber.StatusCreated, map[string]string{"id": id})

}
