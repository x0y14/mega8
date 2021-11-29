package metall

import (
	"github.com/x0y14/mega8/assets/characters"
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/game/compute"
	"github.com/x0y14/mega8/pkg/game/display"
)

func NewIdleMotion() *entity.Motion {
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

	idleAnimation := display.NewAnimation(
		frames,
		compute.Left,
	)

	IdleMotion := entity.NewMotion(
		"idle",
		entity.Idle,
		idleAnimation,
		nil,
	)

	return IdleMotion
}
