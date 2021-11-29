package entity

import (
	"github.com/x0y14/mega8/pkg/game/compute"
	"github.com/x0y14/mega8/pkg/game/display"
)

type Motion struct {
	Name string
	Kind MotionKind
	*display.Animation
	Effects        []compute.Effect
	AllowInterrupt bool
}

func NewMotion(name string, kind MotionKind, animation *display.Animation, effects []compute.Effect, allowInterrupt bool) *Motion {
	return &Motion{
		Name:           name,
		Kind:           kind,
		Animation:      animation,
		Effects:        effects,
		AllowInterrupt: allowInterrupt,
	}
}

func (m *Motion) Update() {
	m.Animation.Update()
}
