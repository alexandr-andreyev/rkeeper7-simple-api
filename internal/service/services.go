package services

import (
	"rkeeper7-simpleapi-service/internal/repository"

	rk7client "github.com/alexandr-andreyev/rk7-xml-apiclient"
)

type IRKeeperService interface {
	GetSystemInfo() (*rk7client.RK7QueryResult, error)
	GetCategList() ([]rk7client.RK7Item, error)
	GetOrderMenu() (*rk7client.RK7QueryResult, error)
	GetMenuItems(refName string, priceType int) (*map[string]interface{}, error)
}

type Services struct {
	RKeeperService IRKeeperService
}

func NewServices(Repos *repository.Repositories) *Services {
	rkeeperService := NewRKeeperService(Repos.RKeeperRepository)

	return &Services{
		RKeeperService: rkeeperService,
	}
}
