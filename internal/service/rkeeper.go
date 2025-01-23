package services

import (
	"rkeeper7-simpleapi-service/internal/repository"

	rk7client "github.com/alexandr-andreyev/rk7-xml-apiclient"
)

type RKeeperService struct {
	repo repository.IRKeeperRepository
}

func NewRKeeperService(repo repository.IRKeeperRepository) *RKeeperService {
	return &RKeeperService{
		repo: repo,
	}
}

func (s RKeeperService) GetSystemInfo() (*rk7client.RK7QueryResult, error) {
	req, err := s.repo.GetSystemInfo()
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (s RKeeperService) GetCategList() ([]rk7client.RK7Item, error) {
	req, err := s.repo.GetCategList()
	if err != nil {
		return nil, err
	}
	items := req.CommandResult[0].Data[0].(rk7client.RK7Reference).Items
	return items, nil
}

func (s RKeeperService) GetMenuItems(refName string, priceType int) (*map[string]interface{}, error) {
	req, err := s.repo.GetRefDataMenuItems()
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	ref := req.CommandResult[0].Data[0].(rk7client.RK7Reference)
	result["products"] = ref.Items
	return &result, nil
}

func (s RKeeperService) GetRestaurants() (*rk7client.RK7QueryResult, error) {
	req, err := s.repo.GetRestaurants()
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (s RKeeperService) GetOrderMenu() (*rk7client.RK7QueryResult, error) {
	req, err := s.repo.GetOrderMenu()
	if err != nil {
		return nil, err
	}

	return req, nil
}
