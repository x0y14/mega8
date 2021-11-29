package metall

import (
	"github.com/x0y14/mega8/assets/characters"
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/game/compute"
	"github.com/x0y14/mega8/pkg/game/display"
)

func NewAttackMotion() *entity.Motion {
	const lifetime = 8
	var sheet = characters.MetallSpritesSheet

	/* frames */
	frame1 := display.NewFrame(
		sheet,
		compute.NewCoordinate(102, 92),
		20,
		23,
		lifetime,
	)

	frame2 := display.NewFrame(
		sheet,
		compute.NewCoordinate(134, 91),
		22,
		24,
		lifetime,
	)

	frame3 := display.NewFrame(
		sheet,
		compute.NewCoordinate(166, 91),
		23,
		24,
		lifetime,
	)

	frame4 := display.NewFrame(
		sheet,
		compute.NewCoordinate(198, 92),
		22,
		25,
		lifetime,
	)

	frame5 := display.NewFrame(
		sheet,
		compute.NewCoordinate(229, 93),
		21,
		25,
		lifetime,
	)

	frame6 := display.NewFrame(
		sheet,
		compute.NewCoordinate(259, 92),
		25,
		25,
		lifetime,
	)

	frame7 := display.NewFrame(
		sheet,
		compute.NewCoordinate(288, 91),
		27,
		26,
		lifetime,
	)

	frames := []display.Frame{
		*frame1,
		*frame2,
		*frame3,
		*frame4,
		*frame5,
		*frame6,
		*frame7,
	}

	attackAnimation := display.NewAnimation(
		frames,
		compute.Left,
	)

	attackMotion := entity.NewMotion(
		"attack",
		entity.Attack,
		attackAnimation,
		nil,
		false,
	)

	return attackMotion
}
