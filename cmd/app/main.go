package main

import (
	"github.com/alxmrz/life/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(internal.NewMainGame(internal.NewGame())); err != nil {
		log.Fatal(err)
	}
}