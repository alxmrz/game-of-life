package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type FakeImage struct {
	ebiten.Image
	isFilled bool
}

func (i *FakeImage) Fill(_ color.Color) {
	i.isFilled = true
}

func (i *FakeImage) IsFilled() bool {
	return i.isFilled
}
