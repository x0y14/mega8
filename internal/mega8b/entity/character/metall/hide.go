package metall

import (
	"github.com/x0y14/mega8/assets/characters"
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/game/compute"
	"github.com/x0y14/mega8/pkg/game/display"
)

func NewHideMotion() *entity.Motion {
	const lifetime = 1
	var sheet = characters.MetallSpritesSheet

	frame1 := display.NewFrame(
		sheet,
		compute.NewCoordinate(176, 133),
		20,
		15,
		lifetime,
	)

	frames := []display.Frame{
		*frame1,
	}

	hideAnimation := display.NewAnimation(
		frames,
		compute.Left,
	)

	hideMotion := entity.NewMotion(
		"hide",
		entity.Hide,
		hideAnimation,
		nil,
		true,
	)

	return hideMotion
}
