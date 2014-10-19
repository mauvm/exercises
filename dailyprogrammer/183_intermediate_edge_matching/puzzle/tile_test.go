package puzzle

import "testing"
import "github.com/stretchr/testify/assert"

func TestTileCreation(t *testing.T) {
	var err error

	tile := CreateTile(862, 0)
	assert.Equal(t, tile.ToString(), "CYMk")

	tile, err = TileFromString("cYkY")
	assert.Equal(t, "cYkY", tile.ToString())

	tile, err = TileFromString("yKMc")
	assert.Equal(t, "yKMc", tile.ToString())

	tile, err = TileFromString("AAAA")
	assert.NotEqual(t, nil, err)
}

func TestRuneToTileValue(t *testing.T) {
	for p, v := range TileValues {
		val, _ := RuneToTileValue(rune(v))
		assert.Equal(t, p, val)
	}
}

func TestTileValues(t *testing.T) {
	tile, _ := TileFromString("KcyM")
	runes := []rune{'K', 'c', 'y', 'M'}

	for p, r := range runes {
		val := tile.Value(uint8(p))
		assert.Equal(t, r, TileValueToRune(val))
	}
}

func TestTileRotating(t *testing.T) {
	var tile Tile
	var i uint8

	for i = 0; i < 4; i += 1 {
		tile = CreateTile(0, i);
		assert.Equal(t, i, tile.GetRotation())
	}

	tile, _ = TileFromString("CYMk")

	assert.Equal(t, 0, tile.GetRotation())
	assert.Equal(t, "N", tile.GetDirection())

	tile.Rotate(1) // 1 step clockwise
	assert.Equal(t, "kCYM", tile.ToString())
	assert.Equal(t, 1, tile.GetRotation())
	assert.Equal(t, "E", tile.GetDirection())

	tile.Rotate(-2) // 2 steps counter clockwise
	assert.Equal(t, "YMkC", tile.ToString())
	assert.Equal(t, 3, tile.GetRotation())
	assert.Equal(t, "W", tile.GetDirection())

	tile.Rotate(7) // 7 % 4 = 3 steps clockwise
	assert.Equal(t, "MkCY", tile.ToString())
	assert.Equal(t, 2, tile.GetRotation())
	assert.Equal(t, "S", tile.GetDirection())
}
