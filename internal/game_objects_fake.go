package internal

type FakeGameObjects struct {
	isUpdated bool
}

func (g *FakeGameObjects) Update() error {
	g.isUpdated = true

	return nil
}

func (g *FakeGameObjects) IsUpdated() bool {
	return g.isUpdated
}

func (g *FakeGameObjects) GetCells() [Width][Height]int {
	return [Width][Height]int{}
}

func (g *FakeGameObjects) Init() error {
	return nil
}
