package liftmanager

import (
	lift "elevatorSystem/src/services/lift"
	"fmt"
	"sort"
)

func (lm *LiftManager) Init(floors int, liftcount int) {
	lm.floors = floors
	lm.lifts = make([]*lift.Lift, liftcount)
	for i := 0; i < liftcount; i++ {
		lm.lifts[i] = lift.NewLift(i)
	}
}

func (lm *LiftManager) RequestNewLift(start int, dest int) int {
	l := -1
	req := NewRequest(start, dest)
	mxTime := 100000

	for i := range lm.lifts {
		if !lm.IsEligible(lm.lifts[i], req) {
			continue
		}
		time := lm.TimeToReach(lm.lifts[i], start)

		if time < mxTime || (time == mxTime && (l == -1 || lm.lifts[i].Id < l)) {
			mxTime = time
			l = i
		}
	}

	if l != -1 {
		lm.lifts[l].AddRequest(start, dest)
	}

	return l
}

func (lm *LiftManager) IsEligible(li *lift.Lift, req Request) bool {
	if li.IsFull() {
		return false
	}
	if li.Direction == "S" {
		return true
	}

	if req.direction == "U" {
		if li.Direction == "D" || li.CurrFloor > req.startFloor {
			return false
		}
		maxStop := maxMapKey(li.UpStops, false)
		if maxStop == -1 {
			maxStop = li.CurrFloor
		}
		if req.startFloor > maxStop {
			return false
		}
	}
	if req.direction == "D" {
		if li.Direction == "U" || li.CurrFloor < req.startFloor {
			return false
		}
		maxStop := maxMapKey(li.DownStops, true)
		if maxStop == -1 {
			maxStop = li.CurrFloor
		}
		if req.startFloor < maxStop {
			return false
		}
	}
	return true
}

func (lm *LiftManager) TimeToReach(li *lift.Lift, floor int) int {
	if li.Direction == "S" {
		return abs(li.CurrFloor - floor)
	}
	time := 0
	pos := li.CurrFloor
	var stops []int
	if li.Direction == "U" {
		stops = sortedKeys(li.UpStops, false)
	} else {
		stops = sortedKeys(li.DownStops, true)
	}
	for i := range stops {
		time += abs(pos - stops[i])
		pos = stops[i]
		if pos == floor {
			return time
		}
	}
	time += abs(pos - floor)
	return time
}

func (lm *LiftManager) Tick() {
	for i := range lm.lifts {
		lm.lifts[i].MoveOneTick()
	}
}

func (em *LiftManager) GetLiftStates() []string {
	states := []string{}
	for _, lift := range em.lifts {
		state := fmt.Sprintf("%d-%s", lift.CurrFloor, lift.Direction)
		states = append(states, state)
	}
	return states
}

func (em *LiftManager) GetNumberOfPeopleOnLift(liftId int) int {
	if liftId < 0 || liftId >= len(em.lifts) {
		return -1
	}
	return em.lifts[liftId].People
}

func (em *LiftManager) GetLiftsStoppingOnFloor(floor int, dir rune) []int {
	ans := []int{}
	for _, lift := range em.lifts {
		if lift.WillStopAt(floor, dir) {
			ans = append(ans, lift.Id)
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

func maxMapKey(m map[int]bool, findMin bool) int {
	best := -1
	for k := range m {
		if best == -1 || (findMin && k < best) || (!findMin && k > best) {
			best = k
		}
	}
	return best
}

func sortedKeys(m map[int]bool, desc bool) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	if desc {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	} else {
		sort.Ints(keys)
	}
	return keys
}
