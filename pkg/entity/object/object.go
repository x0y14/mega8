package object

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/game/compute"
)

type Object struct {
	count int
}

func NewObject() *Object {
	return &Object{}
}

func (o *Object) Draw(screen *ebiten.Image, offsetAdjust bool) {}

func (o *Object) Update() {}

func (o *Object) IsCurrentDirection(direction compute.Direction) bool {
	return false
}

func (o *Object) UpdateDirection() {}

func (o *Object) IsCurrentMotionKind(kind entity.MotionKind) bool {
	return false
}
func (o *Object) UpdateMotionKind() {}

func (o *Object) ListenMotionKindUpdate() {}
