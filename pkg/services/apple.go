package services

import (
	iManagers "example-di/pkg/interfaces/managers"
	"fmt"
)

type AppleService struct {
	manager iManagers.IMainManager
}

func NewAppleService(
	manager iManagers.IMainManager,
) *AppleService {
	return &AppleService{
		manager,
	}
}

func (service *AppleService) A1() {
	fmt.Println("before A1")
	service.manager.Services().BananaService().B1()     // A1 -> B1 -> P1 -> A2
	service.manager.Services().BananaService().B2()     // A1 -> B2 -> A2
	service.manager.Validators().RedValidator().IsRed() // A1 -> IsRed -> A3
	// todo: service.manager.Repos().CustomerRepo().FindByID(123)
	fmt.Println("after A1")
}

func (service *AppleService) A2() {
	fmt.Println("done with A2")
}

func (service *AppleService) A3() {
	fmt.Println("in A3 after color IsRed")
}
