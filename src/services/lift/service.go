package lift

func (l *Lift) IsFull() bool {
	return l.People >= l.MaxPeople
}

func (l *Lift) AddRequest(start, dest int) {
	if start < dest {
		l.UpStops[dest] = true
		l.PickupStops[start] = true
	} else {
		l.DownStops[dest] = true
		l.PickupStops[start] = true
	}
}

func (l *Lift) MoveOneTick() {
	switch l.Direction {
	case "U":
		l.CurrFloor++
		if l.PickupStops[l.CurrFloor] {
			delete(l.PickupStops, l.CurrFloor)
			l.People++
		}
		if l.UpStops[l.CurrFloor] {
			delete(l.UpStops, l.CurrFloor)
			l.People = max(0, l.People-1)
		}
		if len(l.UpStops) == 0 && len(l.PickupStops) == 0 {
			if len(l.DownStops) > 0 {
				l.Direction = "D"
			} else {
				l.Direction = "S"
			}
		}
	case "D":
		l.CurrFloor--
		if l.PickupStops[l.CurrFloor] {
			delete(l.PickupStops, l.CurrFloor)
			l.People++
		}
		if l.DownStops[l.CurrFloor] {
			delete(l.DownStops, l.CurrFloor)
			l.People = max(0, l.People-1)
		}
		if len(l.DownStops) == 0 && len(l.PickupStops) == 0 {
			if len(l.UpStops) > 0 {
				l.Direction = "U"
			} else {
				l.Direction = "S"
			}
		}
	case "S":
		if len(l.UpStops) > 0 || len(l.PickupStops) > 0 {
			l.Direction = "U"
		} else if len(l.DownStops) > 0 {
			l.Direction = "D"
		}
	}
}

func (l *Lift) WillStopAt(floor int, dir rune) bool {
	switch dir {
	case 'U':
		return l.UpStops[floor]
	case 'D':
		return l.DownStops[floor]
	case 'I':
		return l.Direction == "S" && l.CurrFloor == floor
	default:
		return false
	}
}
