package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) CheckKeyBoard() {
	for _, key := range g.keys {
		if key == ebiten.KeyRight {
		} else if key == ebiten.KeyLeft {
		}
	}
}
