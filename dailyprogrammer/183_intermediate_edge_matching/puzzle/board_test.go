package puzzle

import "testing"
import "github.com/stretchr/testify/assert"

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
