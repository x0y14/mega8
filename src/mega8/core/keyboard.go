package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/src/mega8/entities"
)

func (g *Game) CheckKeyBoard() {
	for _, key := range g.keys {
		if key == ebiten.KeyRight {
			g.player.SetDirection(entities.Right)
		} else if key == ebiten.KeyLeft {
			g.player.SetDirection(entities.Left)
		}
	}
}
