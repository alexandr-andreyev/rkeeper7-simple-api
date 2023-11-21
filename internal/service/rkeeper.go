package services

import "fmt"

type RKeeperService struct {
	//repo repository.RKeeper
}

func NewRKeeperService() *RKeeperService {
	return &RKeeperService{
		//repo: repo,
	}
}

func (s RKeeperService) GetSystemInfo() error {
	fmt.Println("TEST GET SYSTEM INFO")
	return nil
}
