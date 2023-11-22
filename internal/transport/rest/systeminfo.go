package rest

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) getSystemInfo(c *fiber.Ctx) error {
	result, err := h.services.RKeeperService.GetSystemInfo()
	if err != nil {
		return err
	}
	return c.JSON(result)
}
