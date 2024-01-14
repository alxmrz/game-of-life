package internal

import (
	"math/rand"
	"time"
)

type GameObjects struct {
	Cells [Width][Height]int
}

func NewGameObjects() *GameObjects {
	return &GameObjects{}
}

func (g *GameObjects) Update() error {
	nextCells := [Width][Height]int{}

	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			alive := 0

			//upper line
			alive += g.getCellByCoordinates(i-1, j-1)
			alive += g.getCellByCoordinates(i, j-1)
			alive += g.getCellByCoordinates(i+1, j-1)

			//middle line except current cell (i,j)
			alive += g.getCellByCoordinates(i-1, j)
			alive += g.getCellByCoordinates(i+1, j)

			//bottom line
			alive += g.getCellByCoordinates(i-1, j+1)
			alive += g.getCellByCoordinates(i, j+1)
			alive += g.getCellByCoordinates(i+1, j+1)

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

func (g *GameObjects) getCellByCoordinates(x, y int) int {
	x, y = g.normalizeCoordinates(x, y)

	return g.Cells[x][y]
}

// Look for out of bounds coordinates and reverse them to opposite position
func (g *GameObjects) normalizeCoordinates(x, y int) (int, int) {
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

func (g *GameObjects) GetCells() [Width][Height]int {
	return g.Cells
}

func (g *GameObjects) Init() error {
	alive := 0
	rand.Seed(time.Now().Unix())
	for i := 0; i < Width; i++ {
		g.Cells[i] = [Height]int{}
		for j := 0; j < Height; j++ {
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
