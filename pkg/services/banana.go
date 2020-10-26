package services

import (
	iManagers "example-di/pkg/interfaces/managers"
	"fmt"
)

type BananaService struct {
	manager iManagers.IMainManager
}

func NewBananaService(
	manager iManagers.IMainManager,
) *BananaService {
	return &BananaService{
		manager,
	}
}

func (service *BananaService) B1() {
	fmt.Println("before B1")
	service.manager.Services().PeachService().P1()
	service.manager.Services().AppleService().A2()
	fmt.Println("after B1")
}

func (service *BananaService) B2() {
	service.manager.Services().AppleService().A2()
}

func (service *BananaService) B3() {
}
