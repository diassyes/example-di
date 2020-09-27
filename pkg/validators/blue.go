package validators

import (
	iManagers "example-di/pkg/interfaces/managers"
)

type Blue struct {
	manager iManagers.IMainManager
}

func NewBlue(
	manager iManagers.IMainManager,
) *Blue {
	return &Blue{
		manager,
	}
}

func (color *Blue) IsBlue() {
}
