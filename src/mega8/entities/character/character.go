package character

import "github.com/x0y14/mega8/src/mega8/entities"

type Character struct {
	OffsetX int
	OffsetY int
	Count   int

	entities.Direction

	MotionState entities.MotionKind
	entities.MotionSet
}

func (c *Character) NowMotion() *entities.Motion {
	return c.MotionSet[c.MotionState]
}

func (c *Character) Update() {
	c.Count++
	c.NowMotion().Update()
}

func (c *Character) SetDirection(direction entities.Direction) {
	c.Direction = direction
}
