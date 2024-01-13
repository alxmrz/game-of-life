package internal

const WIDTH int = 100
const HEIGHT int = 100

type IGame interface {
	Update() error
	Draw(screen ImageInterface) error
	Layout(int, int) (int, int)
	GetCells() [WIDTH][HEIGHT]int
	Init() error
}
