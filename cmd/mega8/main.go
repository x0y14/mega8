package main

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/src/mega8"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 200
	screenHeight = 200
)

func main() {
	ebiten.SetWindowSize(screenWidth+50, screenHeight+50)
	ebiten.SetWindowTitle("mega8")
	gameMega8 := mega8.NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(gameMega8); err != nil {
		log.Fatal(err)
	}
}
