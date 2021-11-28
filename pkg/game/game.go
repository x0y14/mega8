package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/x0y14/mega8/pkg/entity"
)

type Game struct {
	screenWidth  int
	screenHeight int
	count        int
	keys         []ebiten.Key

	Player   entity.Entity
	Entities []entity.Entity
}

func NewGame(screenWidth, screenHeight int) *Game {
	return &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		count:        0,
		keys:         []ebiten.Key{},
	}
}

func (g *Game) SetPlayer(entity *entity.Entity) {
	g.Player = *entity
}

func (g *Game) Update() error {
	g.count++
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	g.Player.Update()
	for _, ent := range g.Entities {
		ent.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen, true)
	for _, ent := range g.Entities {
		ent.Draw(screen, true)
	}

	g.DrawFrameCount(screen)
	g.DrawKeysPressed(screen)

}

func (g *Game) Layout(_, _ int) (int, int) {
	return g.screenWidth, g.screenHeight
}
