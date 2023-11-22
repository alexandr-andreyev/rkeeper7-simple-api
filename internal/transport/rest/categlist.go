package rest

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) getCategList(c *fiber.Ctx) error {
	result, err := h.services.RKeeperService.GetCategList()
	if err != nil {
		return err
	}
	return c.JSON(result)
}
