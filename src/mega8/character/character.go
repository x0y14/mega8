package character

type Character struct {
	OffsetX int
	OffsetY int
	Count   int

	MotionState MotionKind
	MotionSet
}

func (c *Character) Update() {
	c.Count++
	c.NowMotion().Update()
}

func (c *Character) NowMotion() *Motion {
	return c.MotionSet[c.MotionState]
}
