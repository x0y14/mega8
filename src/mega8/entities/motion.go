package entities

import "github.com/x0y14/mega8/src/mega8/entities/work"

type Motion struct {
	Name string
	Kind MotionKind
	*Animation
	Works []work.Work
}

func (m *Motion) Do() {}
