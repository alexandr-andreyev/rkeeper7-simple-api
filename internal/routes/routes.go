package routes

import (
	"rkeeper7-simpleapi-service/internal/transport/rest"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/systeminfo", rest.GetInfo)
	app.Post("/api/v1/refdata", rest.GetRefData)
}
