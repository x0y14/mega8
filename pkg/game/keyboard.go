package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/pkg/game/compute"
)

func (g *Game) CheckKeyBoard() {
	if len(g.keys) == 0 && g.Player.IsMotionEnded() {
		g.Player.DoIdle()
	}

	for _, key := range g.keys {
		if key == ebiten.KeyRight && g.Player.IsCurrentMotionAllowInterrupt() {
			g.Player.DoWalk(compute.Right)
		} else if key == ebiten.KeyLeft {
			g.Player.DoWalk(compute.Left)
		} else if key == ebiten.KeyDown {
			g.Player.DoHide()
		} else if key == ebiten.KeyA {
			g.Player.DoAttack()
		}
	}
}
