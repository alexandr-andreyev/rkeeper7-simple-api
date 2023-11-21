package routes

import (
	"rkeeper7-simpleapi-service/internal/transport/rest"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api/v1")
	//Системная информация о ресторане
	api.Get("/systeminfo", rest.GetInfo)
	//Справочники меню (блюда, цены)
	api.Post("/refdata", rest.GetMenuItems)
	//Справочники группы меню
	api.Post("/categlist", rest.GetCategList)
	//Доступное меню на кассе
	api.Post("/getordermenu", rest.GetOrderMenu)
	// Стоп листы
	api.Get("/dishrests", rest.GetDishRests)
}
