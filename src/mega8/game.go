package mega8

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/src/debugger"
	"github.com/x0y14/mega8/src/mega8/character"
	"github.com/x0y14/mega8/src/mega8/character/metall"
	"github.com/x0y14/mega8/src/mega8/projector"
)

type Game struct {
	screenWidth  int
	screenHeight int
	keys         []ebiten.Key
	count        int
	entities     []character.Character
}

func NewGame(screenWidth, screenHeight int) *Game {

	metall1 := metall.New()

	return &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		keys:         []ebiten.Key{},
		count:        0,
		entities:     []character.Character{*metall1},
	}
}

func (g *Game) Update() error {
	g.count++
	for _, entity := range g.entities {
		entity.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, entity := range g.entities {
		projector.DrawCharacter(screen, entity)
	}
	debugger.DrawFrameCount(g.count, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
