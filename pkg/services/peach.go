package services

import (
	iManagers "example-di/pkg/interfaces/managers"
	"fmt"
)

type PeachService struct {
	manager iManagers.IMainManager
}

func NewPeachService(
	manager iManagers.IMainManager,
) *PeachService {
	return &PeachService{
		manager,
	}
}

func (service *PeachService) P1() {
	fmt.Println("before P1")
	service.manager.Services().AppleService().A2()
	fmt.Println("after P1")
}

func (service *PeachService) P2() {
}

func (service *PeachService) P3() {
}
