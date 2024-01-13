package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainGame_Draw(t *testing.T) {

	game := FakeGame{}

	mainGame := NewMainGame(&game)

	mainGame.Draw(&ebiten.Image{})

	assert.True(t, game.IsDrawn())
}
