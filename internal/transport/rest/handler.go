package rest

import (
	services "rkeeper7-simpleapi-service/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Handler struct {
	services *services.Services
	//tokenManager auth.TokenManager
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
		//tokenManager: tokenManager,
	}
}

func (h Handler) Init() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "RK7 Simple Api v0.0.1",
		DisableStartupMessage: true,
	})

	app.Use(logger.New())

	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	h.setupRoutes(app)

	return app
}

func (h Handler) setupRoutes(app *fiber.App) {
	api := app.Group("api/v1")
	//Системная информация о ресторане
	api.Get("/systeminfo", h.getSystemInfo)
	// //Справочники меню (блюда, цены)
	// api.Post("/refdata", rest.GetMenuItems)
	// //Справочники группы меню
	// api.Post("/categlist", rest.GetCategList)
	// //Доступное меню на кассе
	// api.Post("/getordermenu", rest.GetOrderMenu)
	// // Стоп листы
	// api.Get("/dishrests", rest.GetDishRests)
}
