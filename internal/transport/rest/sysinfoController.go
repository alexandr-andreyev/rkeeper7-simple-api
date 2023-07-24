package rest

import (
	"fmt"
	"rkeeper7-simpleapi-service/internal/config"
	XmlRkeeper "rkeeper7-simpleapi-service/internal/xmlRkeeper"

	"github.com/gofiber/fiber/v2"
)

func GetInfo(c *fiber.Ctx) error {
	req := HttpRequestData{
		CashServerIP:   config.ServerConfig.Config.RKeeperCashServerIp,
		CashServerPort: config.ServerConfig.Config.RkeeperCashServerPort,
		Method:         "POST",
		Username:       config.ServerConfig.Config.RkeeperUser,
		Password:       config.ServerConfig.Config.RkeeperPassword,
	}
	// Request xml body
	xmlData, err := XmlRkeeper.RequestGetSystemInfo()
	if err != nil {
		return err
	}
	req.Payload = xmlData
	//Send http request to Rkeeper7
	body, err := HttpRequestToRkeeper(req)

	if err != nil {
		fmt.Println("err:", err)
		return c.JSON(
			&fiber.Map{
				"data":  body,
				"error": err.Error(),
			},
		)
	}
	resp, err := XmlRkeeper.ResponseGetSystemInfo(body)
	if err != nil {
		fmt.Println("err:", err)
		return c.JSON(
			&fiber.Map{
				"data":  body,
				"error": err.Error(),
			},
		)
	}
	// Обработка ответа
	return c.JSON(fiber.Map{
		"data": resp,
	})
}
