package repository

import rk7client "github.com/alexandr-andreyev/rk7-xml-apiclient"

type rkeeperRepo struct {
}

func NewRkeeperRepo() *rkeeperRepo {
	return &rkeeperRepo{}
}

func (r rkeeperRepo) GetSystemInfo() (*rk7client.RK7QueryResult, error) {
	Client := rk7client.NewClient("127.0.0.1", 14450, "test", "test")

	req, err := Client.GetSystemInfo()
	if err != nil {
		return nil, err
	}

	return req, nil
}
