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

func (s RKeeperService) GetCategList() (*rk7client.RK7QueryResult, error) {
	req, err := s.repo.GetCategList()
	if err != nil {
		return nil, err
	}

	return req, nil
}
