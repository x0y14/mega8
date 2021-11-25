package character

import "github.com/hajimehoshi/ebiten/v2"

type Frame struct {
	Sheet    *ebiten.Image
	Num      int
	OriginX  int
	OriginY  int
	Width    int
	Height   int
	Lifetime int
}
