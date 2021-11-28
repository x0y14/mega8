package entity

type State struct {
	HitPoint  int
	Damage    int
	WalkSpeed int
	RunSpeed  int
}

func NewState() *State {
	return &State{
		HitPoint:  0,
		Damage:    0,
		WalkSpeed: 4,
		RunSpeed:  8,
	}
}
