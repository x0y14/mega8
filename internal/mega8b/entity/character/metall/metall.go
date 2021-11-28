package metall

import (
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/entity/character"
	"github.com/x0y14/mega8/pkg/game/compute"
)

func NewMetall() *entity.Entity {
	motionSet := map[entity.MotionKind]entity.Motion{}
	motionSet[entity.Idle] = *IdleMotion
	motionSet[entity.Walk] = *WalkMotion

	var mtl entity.Entity = character.NewCharacter(
		compute.NewCoordinate(0, 0),
		compute.Left,
		motionSet,
	)
	return &mtl
}
