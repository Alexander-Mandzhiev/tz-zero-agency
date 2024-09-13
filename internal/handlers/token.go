package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetTokenFromHeader(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		err := fmt.Errorf("заголовок авторизации не предоставлен")
		h.log.Error(err.Error())
		return "", err
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		err := fmt.Errorf("неверный формат заголовка авторизации")
		h.log.Error(err.Error())
		return "", err
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	return token, nil
}

func (h *Handler) token(c *fiber.Ctx) (string, error) {
	token, err := h.GetTokenFromHeader(c)
	if err != nil {
		return "", err
	}
	return token, nil
}
