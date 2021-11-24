package mega8

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	screenWidth  int
	screenHeight int
	keys         []ebiten.Key
	count        int
}

func NewGame(screenWidth, screenHeight int) *Game {
	return &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		keys:         []ebiten.Key{},
		count:        0,
	}
}

func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
