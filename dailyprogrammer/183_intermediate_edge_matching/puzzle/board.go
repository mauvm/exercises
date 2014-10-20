package puzzle

import "fmt"
import "errors"

type Board struct {
	size uint8
	tiles []*Tile
}

func CreateBoard(size uint8) Board {
	return Board{size, make([]*Tile, size * size)}
}

func (b Board) GetSize() uint8 {
	return b.size
}

func (b Board) GetTiles() []*Tile {
	return b.tiles
}

func (b Board) GetTile(pos uint8) (t *Tile, err error) {
	if pos >= b.size * b.size {
		return nil, errors.New("Given position is out of range")
	}
	if b.tiles[pos] == nil {
		return nil, errors.New("No tile placed on given position")
	}
	return b.tiles[pos], nil
}

func (b Board) IsSolved() bool {
	var i uint8
	for i = 0; i < b.size; i += 1 {
		if b.tiles[i] == nil || ! b.CanPlace(b.tiles[i], i) {
			return false
		}
	}
	return true
}

func testForAlignedPair(t1 *Tile, t2 *Tile, dir uint8) bool {
	val1 := (*t1).Value(dir)
	val2 := (*t2).Value((dir + 2) % 4)
	if val1 % 2 == 0 {
		// Compare tail to head
		return val1 + 1 == val2
	} else {
		// Compare head to tail
		return val1 - 1 == val2
	}
}

func (b Board) CanPlace(t *Tile, pos uint8) bool {
	if pos >= b.size * b.size {
		return false
	}
	var p_t *Tile
	var err error
	// Test north
	if pos >= b.size {
		if p_t, err = b.GetTile(pos - b.size); err == nil {
			if ! testForAlignedPair(t, p_t, 0) {
				return false
			}
		}
	}
	// Test east
	if pos + 1 % b.size == 0 {
		if p_t, err = b.GetTile(pos + 1); err == nil {
			if ! testForAlignedPair(t, p_t, 1) {
				return false
			}
		}
	}
	// Test south
	if pos < b.size * b.size - b.size {
		if p_t, err = b.GetTile(pos + b.size); err == nil {
			if ! testForAlignedPair(t, p_t, 2) {
				return false
			}
		}
	}
	// Test west
	if pos % b.size > 0 {
		if p_t, err = b.GetTile(pos - 1); err == nil {
			if ! testForAlignedPair(t, p_t, 3) {
				return false
			}
		}
	}
	return true
}

func (b Board) Place(t *Tile, pos uint8) (err error) {
	if ! b.CanPlace(t, pos) {
		message := fmt.Sprintf("Can not place tile %v on position %d", *t, pos)
		return errors.New(message)
	}
	b.tiles[pos] = t
	return nil
}

func (b Board) String() (result string) {
	for _, tile := range b.GetTiles() {
		if tile == nil {
			result += "<nil> "
		} else {
			result += (*tile).String() + " "
		}
	}
	return result
}
