package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) error(c *fiber.Ctx, code int, err error) error {
	h.log.Error(err.Error())
	return h.respond(c, code, map[string]string{"error": err.Error()})
}

func (h *Handler) respond(c *fiber.Ctx, code int, data interface{}) error {
	c.Status(code)
	return c.JSON(data)
}
