package repository

import (
	"rkeeper7-simpleapi-service/internal/config"
	rk7Repo "rkeeper7-simpleapi-service/internal/repository/rkeeper"

	rk7client "github.com/alexandr-andreyev/rk7-xml-apiclient"
)

type IRKeeperRepository interface {
	GetSystemInfo() (*rk7client.RK7QueryResult, error)
	GetCategList() (*rk7client.RK7QueryResult, error)
	GetRefDataMenuItems() (*rk7client.RK7QueryResult, error)
	GetOrderMenu() (*rk7client.RK7QueryResult, error)
}

type Repositories struct {
	RKeeperRepository IRKeeperRepository
}

func NewRepositories(cfg *config.Server) *Repositories {
	rkClient := rk7client.NewClient(
		cfg.Config.RK7ClientConfig.CashServerIp,
		cfg.Config.RK7ClientConfig.CashServerPort,
		cfg.Config.RK7ClientConfig.User,
		cfg.Config.RK7ClientConfig.Password,
	)
	return &Repositories{
		RKeeperRepository: rk7Repo.NewRkeeperRepo(rkClient),
	}
}
