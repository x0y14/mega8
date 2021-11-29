package main

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/internal/mega8b/entity/character/metall"
	"github.com/x0y14/mega8/pkg/game"
	"github.com/x0y14/mega8/pkg/game/compute"
	"log"
)

func main() {
	windowsWidth := 200
	windowsHeight := 200

	mega8b := game.NewGame(windowsHeight, windowsHeight)

	metall1 := *metall.NewMetall("player-metall-1")
	mega8b.SetPlayer(&metall1)

	metall2 := *metall.NewMetall("enemy-metall-1")
	metall2.SetOffset(compute.NewCoordinate(50, 50))
	mega8b.AddEntity(&metall2)

	metall3 := *metall.NewMetall("enemy-metall-2")
	metall3.SetOffset(compute.NewCoordinate(100, 100))
	mega8b.AddEntity(&metall3)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mega8b.ListenEntityMotionUpdate(ctx)

	ebiten.SetWindowSize(windowsWidth*2, windowsHeight*2)
	ebiten.SetWindowTitle("mega8b")

	if err := ebiten.RunGame(mega8b); err != nil {
		log.Fatal(err)
	}
}
