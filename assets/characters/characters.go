package characters

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

//go:embed megaman.png
var MegaManSpritesSheetBytes []byte
var MegaManSpritesSheet *ebiten.Image

//go:embed metall.png
var MetallSpritesSheetBytes []byte
var MetallSpritesSheet *ebiten.Image

func init() {
	// megaman
	buf := bytes.NewReader(MegaManSpritesSheetBytes)
	img, _, err := image.Decode(buf)
	if err != nil {
		log.Fatal(err)
	}
	MegaManSpritesSheet = ebiten.NewImageFromImage(img)

	// metall
	buf = bytes.NewReader(MetallSpritesSheetBytes)
	img, _, err = image.Decode(buf)
	if err != nil {
		log.Fatal(err)
	}
	MetallSpritesSheet = ebiten.NewImageFromImage(img)
}
