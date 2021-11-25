package character

type Motion struct {
	Name string
	Kind MotionKind
	*Animation
}
