package main

import (
	lm "elevatorSystem/src/services/liftManager"
	"fmt"
)

func main() {
	manager := lm.NewLiftManager()
	manager.Init(6, 2)

	fmt.Println(manager.RequestNewLift(0, 3)) // Output: 0
	manager.Tick()
	fmt.Println(manager.GetLiftStates())
	fmt.Println(manager.RequestNewLift(0, 2)) // Output: 1
	manager.Tick()
	fmt.Println(manager.RequestNewLift(0, 5)) // Output: -1
	fmt.Println(manager.RequestNewLift(1, 0)) // Output: 1
	manager.Tick()
	manager.Tick()
	manager.Tick()
}