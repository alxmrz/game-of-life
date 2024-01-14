package main

import (
	"github.com/alxmrz/game-of-life/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game of life")
	if err := ebiten.RunGame(internal.NewMainGame(internal.NewGameObjects())); err != nil {
		log.Fatal(err)
	}
}
