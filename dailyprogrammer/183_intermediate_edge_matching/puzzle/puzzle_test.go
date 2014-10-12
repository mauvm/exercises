package puzzle

import "testing"
import "github.com/stretchr/testify/assert"

// import "fmt"

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
	tile, _ := TileFromString("CYMk")

	assert.Equal(t, "N", TileDirection(tile))

	tile.Rotate(1) // 1 step clockwise
	assert.Equal(t, "kCYM", tile.ToString())
	assert.Equal(t, "E", TileDirection(tile))

	tile.Rotate(-2) // 2 steps counter clockwise
	assert.Equal(t, "YMkC", tile.ToString())
	assert.Equal(t, "W", TileDirection(tile))

	tile.Rotate(7) // 7 % 4 = 3 steps clockwise
	assert.Equal(t, "MkCY", tile.ToString())
	assert.Equal(t, "S", TileDirection(tile))
}

func TestBoardCreation(t *testing.T) {
	board := CreateBoard(3)

	assert.Equal(t, 3, board.size)
	assert.Equal(t, 9, len(board.tiles))
}

func TestBoardTiles(t *testing.T) {
	board := CreateBoard(3)

	var original Tile
	var p_original, p_tile, p_tile_nil *Tile
	var err error

	// Placing and retrieving tile
	original, _ = TileFromString("CYMk")
	p_original = &original
	err = board.Place(p_original, 4)
	assert.Equal(t, nil, err)
	p_tile, err = board.GetTile(4)
	assert.Equal(t, nil, err)
	assert.Equal(t, p_original, p_tile)

	// Yk not possible
	err = board.Place(p_original, 5)
	assert.NotEqual(t, nil, err)

	// Not found
	p_tile, err = board.GetTile(8)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, p_tile_nil, p_tile)

	// Out of range
	p_tile, err = board.GetTile(100)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, p_tile_nil, p_tile)
}

func TestBoardIsSolved(t *testing.T) {
	board := CreateBoard(3)

	assert.Equal(t, false, board.IsSolved())
}
