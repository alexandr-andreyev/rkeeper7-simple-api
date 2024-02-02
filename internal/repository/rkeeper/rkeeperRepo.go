package rk7Repo

import rk7client "rkeeper7-simpleapi-service/pkg/rk7-xml-apiclient"

type rkeeperRepo struct {
	rkClient *rk7client.Client
}

func NewRkeeperRepo(rkClient *rk7client.Client) *rkeeperRepo {
	return &rkeeperRepo{
		rkClient: rkClient,
	}
}

func (r rkeeperRepo) GetSystemInfo() (*rk7client.RK7QueryResult, error) {
	req, err := r.rkClient.GetSystemInfo()
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (r rkeeperRepo) GetCategList() (*rk7client.RK7QueryResult, error) {
	req, err := r.rkClient.GetCateglist()
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (r rkeeperRepo) GetRefDataMenuItems(refName string, priceType int) (*rk7client.RK7QueryResult, error) {
	req, err := r.rkClient.GetRefData(refName, priceType)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (r rkeeperRepo) GetOrderMenu() (*rk7client.RK7QueryResult, error) {
	req, err := r.rkClient.GetOrderMenu("68")
	if err != nil {
		return nil, err
	}

	return req, nil
}
