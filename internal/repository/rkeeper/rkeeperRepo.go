package rk7Repo

import rk7client "github.com/alexandr-andreyev/rk7-xml-apiclient"

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

func (r rkeeperRepo) GetRefDataMenuItems() (*rk7client.RK7QueryResult, error) {
	input := []rk7client.RK7Command{{
		CMD:         rk7client.RK7CMD_GETREFDATA,
		RefName:     rk7client.RK7REF_MENUITEMS,
		OnlyActrive: rk7client.ONLY_ACTIVE_TRUE,
	},
	}
	req, err := r.rkClient.GetRefData(input)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (r rkeeperRepo) GetOrderMenu() (*rk7client.RK7QueryResult, error) {
	req, err := r.rkClient.GetOrderMenu("68") //TODO брать из конфига код станции
	if err != nil {
		return nil, err
	}

	return req, nil
}
