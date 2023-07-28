package rest

import (
	"fmt"
	"rkeeper7-simpleapi-service/internal/config"
	XmlRkeeper "rkeeper7-simpleapi-service/pkg/xmlRkeeper-api"

	"github.com/gofiber/fiber/v2"
)

func GetInfo(c *fiber.Ctx) error {
	Client := XmlRkeeper.NewClient(
		config.ServerConfig.Config.RKeeperCashServerIp,
		config.ServerConfig.Config.RkeeperCashServerPort,
		config.ServerConfig.Config.RkeeperUser,
		config.ServerConfig.Config.RkeeperPassword)

	// Request xml body
	resp, err := Client.GetSystemInfo()

	if err != nil {
		fmt.Println("err:", err.Error())
		return c.JSON(
			&fiber.Map{
				"data":  resp,
				"error": err.Error(),
			},
		)
	}

	// Обработка ответа
	return c.JSON(resp)
}
