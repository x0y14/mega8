package metall

import (
	"github.com/x0y14/mega8/assets/characters"
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/game/compute"
	"github.com/x0y14/mega8/pkg/game/display"
)

var IdleMotion *entity.Motion

func init() {
	const lifetime = 8
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

	IdleMotion = entity.NewMotion(
		"idle",
		entity.Idle,
		idleAnimation,
		nil,
	)
}
