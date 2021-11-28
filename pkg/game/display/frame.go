package display

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/mega8/pkg/game/compute"
	"image"
)

type Frame struct {
	Sheet            *ebiten.Image
	OriginCoordinate *compute.Coordinate
	Width            int
	Height           int
	Lifetime         int
}

func NewFrame(sheet *ebiten.Image, oCoordinate *compute.Coordinate, width, height int, lifetime int) *Frame {
	return &Frame{
		Sheet:            sheet,
		OriginCoordinate: oCoordinate,
		Width:            width,
		Height:           height,
		Lifetime:         lifetime,
	}
}

func (f *Frame) Trim() *ebiten.Image {
	return f.Sheet.SubImage(
		image.Rect(
			f.OriginCoordinate.X, f.OriginCoordinate.Y,
			f.OriginCoordinate.X+f.Width, f.OriginCoordinate.Y+f.Height)).(*ebiten.Image)
}
