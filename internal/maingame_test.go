package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainGame_Update(t *testing.T) {

	gameObjects := FakeGameObjects{}

	mainGame := NewMainGame(&gameObjects)

	mainGame.Update()

	assert.True(t, gameObjects.IsUpdated())
}
