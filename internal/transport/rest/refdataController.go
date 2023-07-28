package rest

import (
	"fmt"
	"rkeeper7-simpleapi-service/internal/config"
	XmlRkeeper "rkeeper7-simpleapi-service/pkg/xmlRkeeper-api"
	"rkeeper7-simpleapi-service/pkg/xmlRkeeper-api/Rk7CMD"

	"github.com/gofiber/fiber/v2"
)

func GetMenuItems(c *fiber.Ctx) error {
	Client := XmlRkeeper.NewClient(
		config.ServerConfig.Config.RKeeperCashServerIp,
		config.ServerConfig.Config.RkeeperCashServerPort,
		config.ServerConfig.Config.RkeeperUser,
		config.ServerConfig.Config.RkeeperPassword)

	// Request xml body
	reqData := new(Rk7CMD.RequestGetRefData)
	reqData.CMD = "GetRefData"
	reqData.RefName = "MenuItems"
	reqData.WithChildItems = "1"

	// Request xml body
	resp, err := Client.GetRefData(*reqData)

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
