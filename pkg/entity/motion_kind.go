package entity

type MotionKind int

const (
	_ MotionKind = iota
	Idle
	Walk
	Run
	Jump
	Attack
	Hurt
	Dead
)

var motionKinds = [...]string{
	Idle:   "Idle",
	Walk:   "Walk",
	Run:    "Run",
	Jump:   "Jump",
	Attack: "Attack",
	Hurt:   "Hurt",
	Dead:   "Dead",
}

func (s MotionKind) String() string {
	return motionKinds[s]
}
