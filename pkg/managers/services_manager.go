package managers

import (
	iManagers "example-di/pkg/interfaces/managers"
	iServices "example-di/pkg/interfaces/services"
	"example-di/pkg/services"
)

type ServicesManager struct {
	appleService  iServices.IAppleService
	bananaService iServices.IBananaService
	peachService  iServices.IPeachService
}

func InitServiceManager(mm iManagers.IMainManager) *ServicesManager {
	return &ServicesManager{
		appleService:  services.NewAppleService(mm),
		bananaService: services.NewBananaService(mm),
		peachService:  services.NewPeachService(mm),
	}
}

func (sm *ServicesManager) AppleService() iServices.IAppleService {
	return sm.appleService
}

func (sm *ServicesManager) BananaService() iServices.IBananaService {
	return sm.bananaService
}

func (sm *ServicesManager) PeachService() iServices.IPeachService {
	return sm.peachService
}
