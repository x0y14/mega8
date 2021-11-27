package metall

import (
	"github.com/x0y14/mega8/assets/characters"
	"github.com/x0y14/mega8/src/mega8/entities"
	"github.com/x0y14/mega8/src/mega8/entities/work"
)

var WalkMotion entities.Motion

func init() {

	const lifetime = 8

	/* frames */
	frame1 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      0,
		OriginX:  4,
		OriginY:  7,
		Width:    22,
		Height:   22,
		Lifetime: lifetime,
	}

	frame2 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      1,
		OriginX:  31,
		OriginY:  7,
		Width:    23,
		Height:   24,
		Lifetime: lifetime,
	}

	frame3 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      2,
		OriginX:  63,
		OriginY:  6,
		Width:    23,
		Height:   25,
		Lifetime: lifetime,
	}

	frame4 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      3,
		OriginX:  95,
		OriginY:  6,
		Width:    23,
		Height:   25,
		Lifetime: lifetime,
	}

	frame5 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      4,
		OriginX:  125,
		OriginY:  5,
		Width:    25,
		Height:   23,
		Lifetime: lifetime,
	}

	frame6 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      5,
		OriginX:  157,
		OriginY:  5,
		Width:    25,
		Height:   24,
		Lifetime: lifetime,
	}

	frame7 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      6,
		OriginX:  189,
		OriginY:  6,
		Width:    26,
		Height:   26,
		Lifetime: lifetime,
	}

	frame8 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      7,
		OriginX:  221,
		OriginY:  7,
		Width:    27,
		Height:   25,
		Lifetime: lifetime,
	}

	frame9 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      8,
		OriginX:  254,
		OriginY:  8,
		Width:    22,
		Height:   24,
		Lifetime: lifetime,
	}

	frame10 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      9,
		OriginX:  283,
		OriginY:  7,
		Width:    21,
		Height:   24,
		Lifetime: lifetime,
	}

	frame11 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      10,
		OriginX:  311,
		OriginY:  6,
		Width:    25,
		Height:   25,
		Lifetime: lifetime,
	}

	frame12 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      11,
		OriginX:  343,
		OriginY:  5,
		Width:    25,
		Height:   26,
		Lifetime: lifetime,
	}

	frame13 := entities.Frame{
		Sheet:    characters.MetallSpritesSheet,
		Num:      12,
		OriginX:  377,
		OriginY:  6,
		Width:    23,
		Height:   25,
		Lifetime: lifetime,
	}

	frames := []entities.Frame{
		frame1,
		frame2,
		frame3,
		frame4,
		frame5,
		frame6,
		frame7,
		frame8,
		frame9,
		frame10,
		frame11,
		frame12,
		frame13,
	}

	/* animation */
	walkAnimation := entities.NewAnimation(
		13,
		frames,
		entities.Left,
		true,
	)

	effectOfWalk := work.Effect{
		EffectKind: work.Move,
	}

	workOfWalk := work.Work{
		EffectRange: nil,
		Effects:     []work.Effect{effectOfWalk},
		Lifetime:    nil,
	}

	/* motion */
	WalkMotion = entities.Motion{
		Name:      "walk",
		Kind:      entities.Walk,
		Animation: walkAnimation,
		Works:     []work.Work{workOfWalk},
	}
}
