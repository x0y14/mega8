package game

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/x0y14/mega8/pkg/entity"
	"github.com/x0y14/mega8/pkg/game/compute"
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

func (g *Game) AddEntity(entity *entity.Entity) {
	g.Entities = append(g.Entities, *entity)
}

func (g *Game) ListenEntityMotionUpdate(ctx context.Context) {
	go g.Player.ListenMotionKindUpdate(ctx)
	for _, ent := range g.Entities {
		go ent.ListenMotionKindUpdate(ctx)
	}
}

func (g *Game) Update() error {
	g.count++
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	g.CheckKeyBoard()

	g.Player.Update()
	for _, ent := range g.Entities {
		ent.Update()
		if g.count%160 == 0 {
			if ent.IsCurrentDirection(compute.Right) {
				ent.UpdateCurrentDirectionIfNotSame(compute.Left)
			} else {
				ent.UpdateCurrentDirectionIfNotSame(compute.Right)
			}
			ent.UpdateCurrentMotionKindIfNotSame(entity.Walk)
		}
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
