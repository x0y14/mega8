package metall

import (
	"github.com/x0y14/mega8/assets/characters"
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/game/compute"
	"github.com/x0y14/mega8/pkg/game/display"
)

func NewWalkMotion() *entity.Motion {
	const lifetime = 8
	var sheet = characters.MetallSpritesSheet

	frame1 := display.NewFrame(
		sheet,
		compute.NewCoordinate(4, 7),
		22,
		22,
		lifetime,
	)

	frame2 := display.NewFrame(
		sheet,
		compute.NewCoordinate(31, 7),
		23,
		24,
		lifetime,
	)

	frame3 := display.NewFrame(
		sheet,
		compute.NewCoordinate(63, 6),
		23,
		25,
		lifetime,
	)

	frame4 := display.NewFrame(
		sheet,
		compute.NewCoordinate(95, 6),
		23,
		25,
		lifetime,
	)

	frame5 := display.NewFrame(
		sheet,
		compute.NewCoordinate(125, 5),
		25,
		23,
		lifetime,
	)

	frame6 := display.NewFrame(
		sheet,
		compute.NewCoordinate(157, 5),
		25,
		24,
		lifetime,
	)

	frame7 := display.NewFrame(
		sheet,
		compute.NewCoordinate(189, 6),
		26,
		26,
		lifetime,
	)

	frame8 := display.NewFrame(
		sheet,
		compute.NewCoordinate(221, 7),
		27,
		25,
		lifetime,
	)

	frame9 := display.NewFrame(
		sheet,
		compute.NewCoordinate(254, 8),
		22,
		24,
		lifetime,
	)

	frame10 := display.NewFrame(
		sheet,
		compute.NewCoordinate(283, 7),
		21,
		24,
		lifetime,
	)

	frame11 := display.NewFrame(
		sheet,
		compute.NewCoordinate(311, 6),
		25,
		25,
		lifetime,
	)

	frame12 := display.NewFrame(
		sheet,
		compute.NewCoordinate(343, 5),
		25,
		26,
		lifetime,
	)

	frame13 := display.NewFrame(
		sheet,
		compute.NewCoordinate(377, 6),
		23,
		25,
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
		*frame8,
		*frame9,
		*frame10,
		*frame11,
		*frame12,
		*frame13,
	}

	walkAnimation := display.NewAnimation(
		frames,
		compute.Left,
	)

	walkMotion := entity.NewMotion(
		"walk",
		entity.Walk,
		walkAnimation,
		nil,
		true,
	)

	return walkMotion
}
