package metall

import (
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/entity/character"
	"github.com/x0y14/mega8/pkg/game/compute"
)

func NewMetall(name string) *entity.Entity {
	motionSet := map[entity.MotionKind]entity.Motion{}
	motionSet[entity.Idle] = *NewIdleMotion()
	motionSet[entity.Walk] = *NewWalkMotion()

	var mtl entity.Entity = character.NewCharacter(
		name,
		compute.NewCoordinate(0, 0),
		compute.Left,
		motionSet,
	)
	return &mtl
}
