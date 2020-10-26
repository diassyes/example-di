package main

import (
	"example-di/pkg/managers"
	"fmt"
)

// 1. No import cycle between packages services/validators.
// 2. No import cycle between services/validators/...:
//    AppleService --> BananaService --> PeachService --> AppleService
// 3. Access to any service/validators from any service/validators bez gemoroya
func main() {
	mainManager := managers.InitMainManager()
	mainManager.Services().AppleService().A1()
	fmt.Println("done")
}
