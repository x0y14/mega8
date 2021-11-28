package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/internal/mega8b/entity/character/metall"
	"github.com/x0y14/mega8/pkg/game"
	"log"
)

func main() {
	windowsWidth := 200
	windowsHeight := 200

	mega8b := game.NewGame(windowsHeight, windowsHeight)
	metall1 := metall.NewMetall()
	mega8b.SetPlayer(metall1)

	ebiten.SetWindowSize(windowsWidth*2, windowsHeight*2)
	ebiten.SetWindowTitle("mega8b")

	if err := ebiten.RunGame(mega8b); err != nil {
		log.Fatal(err)
	}
}
