package managers

type IMainManager interface {
	Services() IServicesManager
	Validators() IValidatorsManager
}
