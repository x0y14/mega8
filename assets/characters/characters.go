package characters

import (
	_ "embed"
	_ "image/png"
)

//go:embed mega-man.png
var MegaManSpritesSheetBytes []byte

//go:embed metall.png
var MetallSpritesSheetBytes []byte
