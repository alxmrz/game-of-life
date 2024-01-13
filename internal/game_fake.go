package internal

type FakeGame struct {
	isDrawn bool
}

func (g *FakeGame) Update() error {
	return nil
}

func (g *FakeGame) Draw(_ ImageInterface) error {
	return nil
}

func (g *FakeGame) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 0, 0
}

func (g *FakeGame) IsDrawn() bool {
	return g.isDrawn
}

func (g *FakeGame) GetCells() [WIDTH][HEIGHT]int {
	return [WIDTH][HEIGHT]int{}
}

func (g *FakeGame) Init() error {
	return nil
}
