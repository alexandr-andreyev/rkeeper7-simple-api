package app

import (
	"fmt"
	"log"

	"rkeeper7-simpleapi-service/internal/config"
	"rkeeper7-simpleapi-service/internal/repository"
	services "rkeeper7-simpleapi-service/internal/service"
	"rkeeper7-simpleapi-service/internal/transport/rest"
)

// The wrapper of your app
func rk7SimpleApiApp(s config.Server) {
	repo := repository.NewRepositories(&s)
	services := services.NewServices(repo)
	handlers := rest.NewHandler(services)
	app := handlers.Init()

	log.Fatal(app.Listen(fmt.Sprintf(":%d", s.Config.ServerPort)))
}
