package repository

import rk7client "github.com/alexandr-andreyev/rk7-xml-apiclient"

type IRKeeperRepository interface {
	GetSystemInfo() (*rk7client.RK7QueryResult, error)
}

type Repositories struct {
	RKeeperRepository IRKeeperRepository
}

func NewRepositories() *Repositories {
	return &Repositories{RKeeperRepository: NewRkeeperRepo()}
}
