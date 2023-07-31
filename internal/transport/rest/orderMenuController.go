package rest

import (
	"fmt"
	"rkeeper7-simpleapi-service/internal/config"
	XmlRkeeper "rkeeper7-simpleapi-service/pkg/xmlRkeeper-api"

	"github.com/gofiber/fiber/v2"
)

func GetOrderMenu(c *fiber.Ctx) error {
	Client := XmlRkeeper.NewClient(
		config.ServerConfig.Config.RKeeperCashServerIp,
		config.ServerConfig.Config.RkeeperCashServerPort,
		config.ServerConfig.Config.RkeeperUser,
		config.ServerConfig.Config.RkeeperPassword)

	// Request xml body

	// Request xml body
	resp, err := Client.GetOrderMenu()

	if err != nil {
		fmt.Println("err:", err)
		return c.JSON(
			&fiber.Map{
				"data":  resp,
				"error": err.Error(),
			},
		)
	}

	return c.JSON(resp)
}
