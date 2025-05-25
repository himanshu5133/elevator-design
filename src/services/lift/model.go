package lift

type Lift struct {
	Id          int
	CurrFloor   int
	Direction   string // 'U', 'D', 'S'
	UpStops     map[int]bool
	DownStops   map[int]bool
	PickupStops map[int]bool
	People      int
	MaxPeople   int
}

func NewLift(id int) *Lift {
	return &Lift{
		Id:          id,
		CurrFloor:   0,
		Direction:   "S",
		UpStops:     make(map[int]bool),
		DownStops:   make(map[int]bool),
		PickupStops: make(map[int]bool),
		MaxPeople:   10,
	}
}
