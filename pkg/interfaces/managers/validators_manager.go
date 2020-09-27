package managers

import iValidators "example-di/pkg/interfaces/validators"

type IValidatorsManager interface {
	RedValidator() iValidators.IRed
	GreenValidator() iValidators.IGreen
	BlueValidator() iValidators.IBlue
}
