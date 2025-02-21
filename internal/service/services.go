package services

import (
	"rkeeper7-simpleapi-service/internal/repository"
	"time"

	rk7client "github.com/alexandr-andreyev/rk7-xml-apiclient"
)

type IRKeeperService interface {
	GetSystemInfo() (*rk7client.RK7QueryResult, error)
	GetCategList() ([]rk7client.RK7Item, error)
	GetOrderMenu() (*rk7client.RK7QueryResult, error)
	GetMenuItems(refName string, priceType int) (*map[string]interface{}, error)
	GetRestaurants() (*rk7client.RK7QueryResult, error)
}

type ICRMService interface {
	GenerateTempCodeFromCard(string, time.Duration) (string, error)
	GetCardByTempCode(tempCode string) (string, error)
}

type Services struct {
	RKeeperService IRKeeperService
	CRMService     ICRMService
}

func NewServices(Repos *repository.Repositories) *Services {
	rkeeperService := NewRKeeperService(Repos.RKeeperRepository)
	crmService := NewCrmService()

	return &Services{
		RKeeperService: rkeeperService,
		CRMService:     crmService,
	}
}
