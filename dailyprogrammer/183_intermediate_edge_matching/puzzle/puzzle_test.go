package puzzle

import "testing"
import "github.com/stretchr/testify/assert"

func TestTileCreation(t *testing.T) {
	var err error

	tile := Tile{862, 0}
	assert.Equal(t, tile.ToString(), "CYMk")

	tile, err = TileFromString("cYkY")
	assert.Equal(t, "cYkY", tile.ToString())

	tile, err = TileFromString("yKMc")
	assert.Equal(t, "yKMc", tile.ToString())

	tile, err = TileFromString("AAAA")
	assert.NotEqual(t, nil, err)
}

func TestRuneToTileValue(t *testing.T) {
	runes := []rune{'c', 'C', 'm', 'M', 'y', 'Y', 'k', 'K'}
	for p, r := range runes {
		val, _ := RuneToTileValue(r)
		assert.Equal(t, p, val)
	}
}

func TestTileValues(t *testing.T) {
	tile, _ := TileFromString("KcyM")
	runes := []rune{'K', 'c', 'y', 'M'}
	values := []uint8{7, 0, 4, 3}

	for p, r := range runes {
		val := TileValue(tile, uint8(p))
		assert.Equal(t, values[p], val)
		assert.Equal(t, r, TileValueToRune(val))
	}
}

func TestTileRotating(t *testing.T) {
	tile := Tile{862, 0} // CYMk

	tile.Rotate(1) // 1 step clockwise
	assert.Equal(t, "kCYM", tile.ToString())

	tile.Rotate(-2) // 2 steps counter clockwise
	assert.Equal(t, "YMkC", tile.ToString())

	tile.Rotate(7) // 7 % 4 = 3 steps clockwise
	assert.Equal(t, "MkCY", tile.ToString())
}
