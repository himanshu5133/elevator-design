package liftmanager

import lift "elevatorSystem/src/services/lift"

type LiftManager struct {
	floors int
	lifts  []*lift.Lift
}

type Request struct {
	startFloor       int
	destinationFloor int
	direction        string // 'U' or 'D' or 'S'
}

func NewLiftManager() *LiftManager {
	return &LiftManager{}
}

func NewRequest(start int, dest int) Request {
	var dir string
	if start < dest {
		dir = "U"
	} else {
		dir = "D"
	}
	return Request{
		startFloor:       start,
		destinationFloor: dest,
		direction:        dir,
	}
}
