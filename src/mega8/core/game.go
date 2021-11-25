package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/src/debugger"
	"github.com/x0y14/mega8/src/mega8/entities/character"
	"github.com/x0y14/mega8/src/mega8/entities/character/metall"
	"github.com/x0y14/mega8/src/mega8/projector"
)

type Game struct {
	screenWidth  int
	screenHeight int
	keys         []ebiten.Key
	count        int
	player       *character.Character
	entities     []character.Character
}

func NewGame(screenWidth, screenHeight int) *Game {

	metall0 := metall.New()
	metall1 := metall.New()

	return &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		keys:         []ebiten.Key{},
		count:        0,
		player:       metall0,
		entities:     []character.Character{*metall1},
	}
}

func (g *Game) Update() error {
	g.count++

	g.player.Update()
	for _, entity := range g.entities {
		entity.Update()
	}
	g.CheckKeyBoard()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, entity := range g.entities {
		projector.DrawCharacter(screen, entity, true)
	}
	debugger.DrawFrameCount(g.count, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
