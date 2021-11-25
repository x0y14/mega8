package character

type Character struct {
	OffsetX int
	OffsetY int
	Count   int

	MotionState MotionKind
	MotionSet
}

func (c *Character) CountUp() {
	c.Count++
}

func (c *Character) NowMotion() *Motion {
	return c.MotionSet[c.MotionState]
}
