package entities

type Situation int

const (
	_ Situation = iota

	Invisible
	Static

	Idling
	Hurting
	Dying

	Hiding

	Walking
	Running

	Jumping
	DoubleJumping
	Falling
	Flying

	Attacking
	RunAttacking
	JumpAttacking
)
