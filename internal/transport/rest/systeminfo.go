package rest

import (
	"github.com/gofiber/fiber/v2"
)

func getSystemInfo(c *fiber.Ctx) error {
	return c.JSON("test")
}
