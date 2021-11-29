package entity

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/pkg/game/compute"
)

type Entity interface {
	Draw(screen *ebiten.Image, offsetAdjust bool)
	Update()

	SetOffset(coordinate *compute.Coordinate)

	IsCurrentDirection(direction compute.Direction) bool
	UpdateCurrentDirection(direction compute.Direction)
	UpdateCurrentDirectionIfNotSame(direction compute.Direction)

	IsCurrentMotionKind(kind MotionKind) bool
	UpdateCurrentMotionKind(kind MotionKind)
	UpdateCurrentMotionKindIfNotSame(kind MotionKind)

	ListenMotionKindUpdate(ctx context.Context)
}
