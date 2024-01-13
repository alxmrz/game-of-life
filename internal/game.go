package internal

import (
	"image/color"
	"math/rand"
	"time"
)

type Game struct {
	Cells [WIDTH][HEIGHT]int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	nextCells := [WIDTH][HEIGHT]int{}

	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			alive := 0

			if g.getCellByIndex(i-1, j-1) == 1 {
				alive++
			}

			if g.getCellByIndex(i, j-1) == 1 {
				alive++
			}

			if g.getCellByIndex(i+1, j-1) == 1 {
				alive++
			}

			if g.getCellByIndex(i-1, j) == 1 {
				alive++
			}

			if g.getCellByIndex(i, j) == 1 {
				//alive++
			}

			if g.getCellByIndex(i+1, j) == 1 {
				alive++
			}

			if g.getCellByIndex(i-1, j+1) == 1 {
				alive++
			}

			if g.getCellByIndex(i, j+1) == 1 {
				alive++
			}
			if g.getCellByIndex(i+1, j+1) == 1 {
				alive++
			}

			if alive > 3 || alive < 2 {
				nextCells[i][j] = 0
			} else if alive == 3 {
				nextCells[i][j] = 1
			} else if alive == 2 {
				nextCells[i][j] = g.GetCells()[i][j]
			}

		}
	}

	g.Cells = nextCells
	return nil
}

func (g *Game) getCellByIndex(x, y int) int {
	x, y = g.normalizeIndex(x, y)

	return g.Cells[x][y]
}

func (g *Game) normalizeIndex(x, y int) (int, int) {
	if x < 0 {
		x = len(g.Cells) + x
	} else if x >= len(g.Cells) {
		x = x - len(g.Cells)
	}

	if y < 0 {
		y = len(g.Cells) + y
	} else if y >= len(g.Cells) {
		y = y - len(g.Cells)
	}
	return x, y
}

func (g *Game) Draw(screen ImageInterface) error {
	screen.Fill(color.RGBA{128, 128, 128, 0})

	return nil
}

func (g *Game) Layout(int, int) (screenWidth, screenHeight int) {
	return 320, 340
}

func (g *Game) GetCells() [WIDTH][HEIGHT]int {
	return g.Cells
}

func (g *Game) Init() error {
	alive := 0
	rand.Seed(time.Now().Unix())
	for i := 0; i < WIDTH; i++ {
		g.Cells[i] = [HEIGHT]int{}
		for j := 0; j < HEIGHT; j++ {
			random := rand.Intn(2)
			if random == 1 {
				if alive == 5000 {
					g.Cells[i][j] = 0
				} else {
					g.Cells[i][j] = 1
					alive++
				}
			} else {
				g.Cells[i][j] = 0
			}
		}
	}

	return nil
}
