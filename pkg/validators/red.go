package validators

import (
	iManagers "example-di/pkg/interfaces/managers"
	"fmt"
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
	fmt.Println("i'm red")
	color.manager.Services().AppleService().A3()
}
