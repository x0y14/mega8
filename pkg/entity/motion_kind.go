package entity

type MotionKind int

const (
	_ MotionKind = iota
	Idle
	Hide
	Walk
	Run
	Jump
	Attack
	Hurt
	Dead
)

var motionKinds = [...]string{
	Idle:   "Idle",
	Hide:   "Hide",
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
