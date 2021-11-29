package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/game/compute"
)

func (g *Game) CheckKeyBoard() {
	if len(g.keys) == 0 {
		g.Player.UpdateCurrentMotionKindIfNotSame(entity.Idle)
	}

	for _, key := range g.keys {
		if key == ebiten.KeyRight {
			g.Player.UpdateCurrentDirectionIfNotSame(compute.Right)
			g.Player.UpdateCurrentMotionKindIfNotSame(entity.Walk)
		} else if key == ebiten.KeyLeft {
			g.Player.UpdateCurrentDirectionIfNotSame(compute.Left)
			g.Player.UpdateCurrentMotionKindIfNotSame(entity.Walk)
		}
	}
}
