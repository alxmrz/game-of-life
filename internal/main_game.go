package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"time"
)

const Width int = 100
const Height int = 100
const CellSize int = 10

type GameObjectsInterface interface {
	Update() error
	GetCells() [Width][Height]int
	Init() error
}

// MainGame implement ebiten.Game interface required by ebiten engine.
type MainGame struct {
	gameObjects GameObjectsInterface
}

func NewMainGame(gameObjects GameObjectsInterface) *MainGame {
	gameObjects.Init()
	mainGame := &MainGame{}

	mainGame.gameObjects = gameObjects

	return mainGame
}

func (g *MainGame) Update() error {
	err := g.gameObjects.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *MainGame) Draw(screen *ebiten.Image) {
	time.Sleep(100 * time.Millisecond)
	screen.Fill(color.RGBA{128, 128, 128, 0})

	cells := g.gameObjects.GetCells()

	for i, line := range cells {
		for j, cell := range line {
			x := float32(i * CellSize)
			y := float32(j * CellSize)
			var clr color.RGBA
			if cell == 1 {
				clr = color.RGBA{255, 0, 0, 0}
			} else {
				clr = color.RGBA{0, 255, 255, 255}
			}
			vector.DrawFilledRect(screen, x, y, float32(CellSize), float32(CellSize), clr, false)
			vector.StrokeRect(screen, x, y, float32(CellSize), float32(CellSize), 1, color.RGBA{0, 0, 0, 0}, false)

		}
	}
}

func (g *MainGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width * CellSize, Height * CellSize
}
