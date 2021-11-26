package metall

import (
	"github.com/x0y14/mega8/src/mega8/entities"
	character2 "github.com/x0y14/mega8/src/mega8/entities/character"
)

type Metall character2.Character

func New() *character2.Character {
	motionSet := entities.MotionSet{
		entities.Walk:   &WalkMotion,
		entities.Attack: &AttackMotion,
	}
	return &character2.Character{
		OffsetX:     0,
		OffsetY:     0,
		Count:       0,
		Direction:   entities.Left,
		MotionState: entities.Walk,
		MotionSet:   motionSet,
	}
}
