package debugger

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

var fontFace font.Face

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	fontFace, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func DrawFrameCount(count int, screen *ebiten.Image) {
	{
		const x, y = 0, 10
		b := text.BoundString(fontFace, fmt.Sprintf("frame: %v", count))
		ebitenutil.DrawRect(screen, float64(b.Min.X+x), float64(b.Min.Y+y), float64(b.Dx()), float64(b.Dy()), color.Transparent)
		text.Draw(screen, fmt.Sprintf("frame: %v", count), fontFace, x, y, color.White)
	}
}
