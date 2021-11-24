package main

import (
	_ "embed"
	"fmt"
	"github.com/x0y14/mega8/assets/characters"
	_ "image/png"
)

func main() {
	fmt.Printf("%v", characters.MegaManSpritesSheetBytes)
	fmt.Printf("%v", characters.MetallSpritesSheetBytes)
}
