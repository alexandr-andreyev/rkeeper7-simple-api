package rest

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetOrderMenu(c *fiber.Ctx) error {
	result, err := h.services.RKeeperService.GetOrderMenu()
	if err != nil {
		return err
	}
	return c.JSON(result)
}
