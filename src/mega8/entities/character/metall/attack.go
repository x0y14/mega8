package metall

import (
	"github.com/x0y14/mega8/assets/characters"
	"github.com/x0y14/mega8/src/mega8/entities"
	"github.com/x0y14/mega8/src/mega8/entities/work"
)

var AttackMotion entities.Motion

func init() {

	const lifetime = 8
	var sheet = characters.MetallSpritesSheet

	/* frames */
	frame1 := entities.Frame{
		Sheet:    sheet,
		Num:      0,
		OriginX:  102,
		OriginY:  92,
		Width:    20,
		Height:   23,
		Lifetime: lifetime,
	}

	frame2 := entities.Frame{
		Sheet:    sheet,
		Num:      1,
		OriginX:  134,
		OriginY:  91,
		Width:    22,
		Height:   24,
		Lifetime: lifetime,
	}

	frame3 := entities.Frame{
		Sheet:    sheet,
		Num:      2,
		OriginX:  166,
		OriginY:  91,
		Width:    23,
		Height:   24,
		Lifetime: lifetime,
	}

	frame4 := entities.Frame{
		Sheet:    sheet,
		Num:      3,
		OriginX:  198,
		OriginY:  92,
		Width:    22,
		Height:   25,
		Lifetime: lifetime,
	}

	frame5 := entities.Frame{
		Sheet:    sheet,
		Num:      4,
		OriginX:  229,
		OriginY:  93,
		Width:    21,
		Height:   25,
		Lifetime: lifetime,
	}

	frame6 := entities.Frame{
		Sheet:    sheet,
		Num:      5,
		OriginX:  259,
		OriginY:  92,
		Width:    25,
		Height:   25,
		Lifetime: lifetime,
	}

	frame7 := entities.Frame{
		Sheet:    sheet,
		Num:      6,
		OriginX:  288,
		OriginY:  91,
		Width:    27,
		Height:   26,
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
	}

	/* animation */
	attackAnimation := entities.Animation{
		Direction:     entities.Left,
		FrameNum:      7,
		NowFrameNo:    0,
		LifetimeCount: 0,
		Frames:        frames,
	}

	effectOfAttack := work.Effect{
		EffectKind: work.Attack,
	}

	workOfAttack := work.Work{
		EffectRange: nil,
		Effects:     []work.Effect{effectOfAttack},
		Lifetime:    nil,
	}

	/* motion */
	AttackMotion = entities.Motion{
		Name:      "walk",
		Kind:      entities.Walk,
		Animation: &attackAnimation,
		Works:     []work.Work{workOfAttack},
	}
}
