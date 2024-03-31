package main

import (
	"hello-world/cat"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	lv := os.Getenv("LOG_LEVEL")

	ebiten.SetWindowSize(cat.ScreenWidth, cat.ScreenHeight)
	ebiten.SetWindowTitle("Cat game")
	if err := ebiten.RunGame(cat.NewGame(lv)); err != nil {
		log.Fatal(err)
	}
}
