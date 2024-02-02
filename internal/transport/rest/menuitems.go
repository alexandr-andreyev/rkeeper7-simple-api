package rest

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetMenuItems(c *fiber.Ctx) error {
	result, err := h.services.RKeeperService.GetMenuItems("MENUITEMS", 3)
	if err != nil {
		return err
	}
	return c.JSON(result)
}
