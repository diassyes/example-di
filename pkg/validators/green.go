package validators

import (
	iManagers "example-di/pkg/interfaces/managers"
)

type Green struct {
	manager iManagers.IMainManager
}

func NewGreen(
	manager iManagers.IMainManager,
) *Green {
	return &Green{
		manager,
	}
}

func (color *Green) IsGreen() {
}
