package puzzle

// cCmMyYkK = 0 1 2 3 4 5 6 7
// tail = cmyk = (x % 2) == 0
// head = CMYK = (x % 2) == 1
// Tile: 0000 0011 0101 1110 = (0000) 001 101 011 110 = (0) 1 5 3 6 = CYMk
// NOTE: This binary setup only allows for 8 card values

import "errors"

var TileValues = []uint8{99, 67, 109, 77, 121, 89, 107, 75}

type Tile struct {
	val uint16
	orientation uint8
}

func (t *Tile) Rotate(steps int) {
	// Modulo that always returns a positive number
	// Also see: https://code.google.com/p/go/issues/detail?id=448
	steps = ((steps % 4) + 4) % 4 // -1 becomes 3

	if steps == 0 {
		return
	}

	// Steps:
	// - Reset trailing bits (with val & 4095)
	// - Shift value to amount of steps to the right
	// - Determine shifted of bits (with (1 << uint8(steps * 3) - 1) & t.val)
	// - Shift those bits one step to the left
	val := t.val
	t.val = (val & 4095) >> uint8(steps * 3)
	t.val |= ((1 << uint8(steps * 3) - 1) & val) << uint8((4 - steps) * 3)
}

func (t Tile) ToString() string {
	val := string(TileValueToRune(TileValue(t, 0)))
	val += string(TileValueToRune(TileValue(t, 1)))
	val += string(TileValueToRune(TileValue(t, 2)))
	val += string(TileValueToRune(TileValue(t, 3)))
	return val
}

func TileFromString(input string) (t Tile, err error) {
	if len(input) != 4 {
		return Tile{0, 0}, errors.New("String length must be 4")
	}
	runes := []rune(input)
	val := uint16(0)
	for i := 0; i < 4; i += 1 {
		tempVal, err := RuneToTileValue(runes[i])
		if err != nil {
			return Tile{0, 0}, errors.New("Invalid character " + string(runes[0]))
		}
		val += uint16(tempVal) << (uint8(3 - i) * 3)
	}
	return Tile{val, 0}, nil
}

func TileValue(t Tile, pos uint8) (val uint8) {
	// Steps:
	// - Modulo with 4 to allow 0-3
	// - Invert position
	// - Multiple with 3 (because 4 sets of 3 bytes)
	// - Binary AND (& 7) to strip off trailing bytes
	return uint8((t.val >> ((3 - (pos % 4)) * 3)) & 7)
}

func TileValueToRune(pos uint8) (r rune) {
	return rune(TileValues[pos % 8])
}

func RuneToTileValue(r rune) (val uint8, err error) {
	for p, v := range TileValues {
		if (v == uint8(r)) {
			return uint8(p), nil;
		}
	}
	return 0, errors.New("No value for rune");
}

type Board []Tile

func CreateBoard(size uint8) Board {
	return Board{}
}

func (b Board) CanPlace(t Tile, pos uint8) bool {
	return true // TODO
}

func (b Board) Place(t Tile, pos uint8) (err error) {
	return nil // TODO
}
