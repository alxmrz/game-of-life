package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"time"
)

// MainGame implement ebiten.Game interface required by ebiten engine.
// Interface main.IGame exist for convenient testing
type MainGame struct {
	game IGame
}

func NewMainGame(game IGame) *MainGame {
	game.Init()
	mainGame := &MainGame{}

	mainGame.game = game

	return mainGame
}

func (g *MainGame) Update() error {
	err := g.game.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *MainGame) Draw(screen *ebiten.Image) {
	time.Sleep(100 * time.Millisecond)

	err := g.game.Draw(screen)
	if err != nil {
		return
	}
	cells := g.game.GetCells()

	for i, line := range cells {
		for j, cell := range line {
			x := float32(i * 10)
			y := float32(j * 10)
			var clr color.RGBA
			if cell == 1 {
				clr = color.RGBA{255, 0, 0, 0}
			} else {
				clr = color.RGBA{0, 255, 255, 255}
			}
			vector.DrawFilledRect(screen, x, y, 10, 10, clr, false)
			vector.StrokeRect(screen, x, y, 10, 10, 1, color.RGBA{0, 0, 0, 0}, false)

		}
	}

	//newImage := ebiten.Image{}
	ebitenutil.DebugPrint(screen, "Debug info")
}

func (g *MainGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1000, 1000
}
