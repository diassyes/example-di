package managers

import (
	iManagers "example-di/pkg/interfaces/managers"
)

type MainManager struct {
	services   iManagers.IServicesManager
	validators iManagers.IValidatorsManager
}

func (mm *MainManager) InitMainManager() {
	mm.services = InitServiceManager(mm)
	mm.validators = InitValidatorsManager(mm)
}

func (mm *MainManager) Services() iManagers.IServicesManager {
	return mm.services
}

func (mm *MainManager) Validators() iManagers.IValidatorsManager {
	return mm.validators
}
