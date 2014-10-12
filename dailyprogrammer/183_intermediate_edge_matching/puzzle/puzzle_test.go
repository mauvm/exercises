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
	for p, v := range TileValues {
		val, _ := RuneToTileValue(rune(v))
		assert.Equal(t, p, val)
	}
}

func TestTileValues(t *testing.T) {
	tile, _ := TileFromString("KcyM")
	runes := []rune{'K', 'c', 'y', 'M'}

	for p, r := range runes {
		val := TileValue(tile, uint8(p))
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
	original, _ = TileFromString("CYck")
	p_original = &original
	err = board.Place(p_original, 4)
	assert.Equal(t, nil, err)
	p_tile, err = board.GetTile(4)
	assert.Equal(t, nil, err)
	assert.Equal(t, p_original, p_tile)

	// Yk is not possible
	err = board.Place(p_original, 5)
	assert.NotEqual(t, nil, err)

	// cC is possible
	err = board.Place(p_original, 1)
	assert.Equal(t, nil, err)

	// Not found
	p_tile, err = board.GetTile(8)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, p_tile_nil, p_tile)

	// Out of range
	p_tile, err = board.GetTile(100)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, p_tile_nil, p_tile)
}

func TestBoardIsNotSolved(t *testing.T) {
	board := CreateBoard(3)

	assert.Equal(t, false, board.IsSolved())

	tile, _ := TileFromString("CYck")
	board.Place(&tile, 4)

	assert.Equal(t, false, board.IsSolved())
}

func TestBoardIsSolved(t *testing.T) {
	board := CreateBoard(3)
	tiles := []string{"CYMk", "McKy", "YMKC", "mCmK", "kyMc", "kYcY", "MKyC", "mYCk", "CMky"}

	for i := 0; i < len(tiles); i += 1 {
		tile, err := TileFromString(tiles[i])
		assert.Equal(t, nil, err)
		err = board.Place(&tile, uint8(i))
		assert.Equal(t, nil, err)
	}

	assert.Equal(t, true, board.IsSolved())
}
