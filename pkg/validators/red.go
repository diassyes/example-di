package validators

import (
	iManagers "example-di/pkg/interfaces/managers"
)

type Red struct {
	manager iManagers.IMainManager
}

func NewRed(
	manager iManagers.IMainManager,
) *Red {
	return &Red{
		manager,
	}
}

func (color *Red) IsRed() {
	color.manager.Services().AppleService().A3()
}
