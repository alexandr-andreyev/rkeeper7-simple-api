package routes

import (
	"rkeeper7-simpleapi-service/internal/transport/rest"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api/v1")
	api.Get("/systeminfo", rest.GetInfo)
	api.Post("/refdata", rest.GetMenuItems)
	api.Post("/categlist", rest.GetCategList)
	api.Post("/getordermenu", rest.GetOrderMenu)
}
