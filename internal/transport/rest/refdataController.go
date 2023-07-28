package rest

import (
	"fmt"
	"rkeeper7-simpleapi-service/internal/config"
	XmlRkeeper "rkeeper7-simpleapi-service/pkg/xmlRkeeper-api"
	XMLRK7Models "rkeeper7-simpleapi-service/pkg/xmlRkeeper-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetRefData(c *fiber.Ctx) error {
	req := HttpRequestData{
		CashServerIP:   config.ServerConfig.Config.RKeeperCashServerIp,
		CashServerPort: config.ServerConfig.Config.RkeeperCashServerPort,
		Method:         "POST",
		Username:       config.ServerConfig.Config.RkeeperUser,
		Password:       config.ServerConfig.Config.RkeeperPassword,
	}
	// Request xml body
	reqData := new(XMLRK7Models.RequestGetRefData)
	reqData.CMD = "GetRefData"
	reqData.RefName = "MenuItems"
	reqData.IgnoreEnums = "1"
	reqData.WithChildItems = "1"
	reqData.WithMacroProp = "1"
	xmlData, err := XmlRkeeper.RequestGetRefData(reqData)
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

	return c.JSON(fiber.Map{"data": string(body)})
}
