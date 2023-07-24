package app

import (
	"fmt"
	"log"

	"rkeeper7-simpleapi-service/internal/config"
	"rkeeper7-simpleapi-service/internal/routes"

	"github.com/gofiber/fiber/v2"
)

// The wrapper of your app
func rk7SimpleApiApp(s config.Server) {
	//TODO Добавить логирование в файл
	s.Winlog.Info(1, "In app.rk7SimpleApi")

	app := fiber.New()
	routes.Setup(app)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", s.Config.ServerPort)))
}
