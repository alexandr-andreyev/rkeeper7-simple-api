package app

import (
	"fmt"
	"log"

	"rkeeper7-simpleapi-service/internal/config"
	services "rkeeper7-simpleapi-service/internal/service"
	"rkeeper7-simpleapi-service/internal/transport/rest"
)

// The wrapper of your app
func rk7SimpleApiApp(s config.Server) {
	services := services.NewServices()
	handlers := rest.NewHandler(services)
	app := handlers.Init()

	log.Fatal(app.Listen(fmt.Sprintf(":%d", s.Config.ServerPort)))
}
