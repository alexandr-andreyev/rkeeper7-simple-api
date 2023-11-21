package services

type IRKeeperService interface {
	GetSystemInfo() error
}

type Services struct {
	RKeeperService IRKeeperService
}

func NewServices() *Services {
	rkeeperService := NewRKeeperService()

	return &Services{
		RKeeperService: rkeeperService,
	}
}
