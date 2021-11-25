package main

import (
	"github.com/x0y14/mega8/src/downloader"
	"os"
)

func main() {
	// mega man
	const megaManSheetPath = "../assets/characters/megaman.png"
	const megaManSpritesSheetUrl = "https://www.spriters-resource.com/download/68847/"
	if _, err := os.Stat(megaManSheetPath); os.IsNotExist(err) {
		if err := downloader.Download(megaManSpritesSheetUrl, megaManSheetPath); err != nil {
			panic(err)
		}
	}

	// metall
	const metallSheetPath = "../assets/characters/metall.png"
	const metallSpritesSheetUrl = "https://www.spriters-resource.com/download/85346/"
	if _, err := os.Stat(metallSheetPath); os.IsNotExist(err) {
		if err := downloader.Download(metallSpritesSheetUrl, metallSheetPath); err != nil {
			panic(err)
		}
	}
}
