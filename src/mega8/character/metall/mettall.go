package metall

import "github.com/x0y14/mega8/src/mega8/character"

type Metall character.Character

func New() *character.Character {
	motionSet := character.MotionSet{
		character.Walk: &WalkMotion,
	}
	return &character.Character{
		OffsetX:     0,
		OffsetY:     0,
		Count:       0,
		MotionState: character.Walk,
		MotionSet:   motionSet,
	}
}
