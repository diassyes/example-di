package managers

import (
	iManagers "example-di/pkg/interfaces/managers"
	iValidators "example-di/pkg/interfaces/validators"
	"example-di/pkg/validators"
)

type ValidatorsManager struct {
	redValidator   iValidators.IRed
	greenValidator iValidators.IGreen
	blueValidator  iValidators.IBlue
}

func InitValidatorsManager(mm iManagers.IMainManager) *ValidatorsManager {
	return &ValidatorsManager{
		redValidator:   validators.NewRed(mm),
		greenValidator: validators.NewGreen(mm),
		blueValidator:  validators.NewBlue(mm),
	}
}

func (vm *ValidatorsManager) RedValidator() iValidators.IRed {
	return vm.redValidator
}

func (vm *ValidatorsManager) GreenValidator() iValidators.IGreen {
	return vm.greenValidator
}

func (vm *ValidatorsManager) BlueValidator() iValidators.IBlue {
	return vm.blueValidator
}
