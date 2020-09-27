package managers

import iServices "example-di/pkg/interfaces/services"

type IServicesManager interface {
	AppleService() iServices.IAppleService
	BananaService() iServices.IBananaService
	PeachService() iServices.IPeachService
}
