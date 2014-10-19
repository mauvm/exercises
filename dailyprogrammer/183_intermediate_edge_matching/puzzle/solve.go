package puzzle

import "errors"

func copyBoard(size uint8, board Board) Board {
	boardCopy := CreateBoard(size)
	copyTiles := boardCopy.GetTiles()
	for i, tile := range board.GetTiles() {
		if tile == nil {
			continue
		}
		newTile := Tile(*tile)
		copyTiles[i] = &newTile
	}
	return boardCopy
}

func copyTilesAndRemoveIndex(tiles []Tile, skip uint8) []Tile {
	newTiles := make([]Tile, len(tiles) - 1)
	j := uint8(0)
	for i, tile := range tiles {
		if uint8(i) == skip {
			continue
		}
		newTiles[j] = tile
		j += 1
	}
	return newTiles
}

func solveStep(size uint8, board Board, pos uint8, tiles []Tile, solutions chan Board) {
	var i uint8
	var j int
	var tile Tile
	tileCount := uint8(len(tiles))
	isLastPos := (pos == size * size - 1)
	for i = 0; i < tileCount; i += 1 {
		for j = 0; j < 4; j += 1 {
			boardCopy := copyBoard(size, board)
			tile = Tile(uint16(tiles[i]))
			if j > 0 {
				tile.Rotate(j)
			}
			err := boardCopy.Place(&tile, pos)
			if err != nil {
				continue
			}
			if isLastPos {
				if boardCopy.IsSolved() {
					solutions <-boardCopy
				}
			} else {
				// TODO: Make fully concurrent
				solveStep(size, boardCopy, pos + 1, copyTilesAndRemoveIndex(tiles, i), solutions)
			}
		}
	}
}

func Solve(size uint8, tiles []Tile) (solutions chan Board, err error) {
	board := CreateBoard(size)
	solutions = make(chan Board)

	if board.GetSize() * board.GetSize() != uint8(len(tiles)) {
		return nil, errors.New("Board size and number of tiles do not match")
	}

	go solveStep(size, board, 0, tiles, solutions)

	return solutions, err
}

func BoardToSolution(board Board, tiles []Tile) (solution string, err error) {
	if ! board.IsSolved() {
		return solution, errors.New("Not solved")
	}

	// TODO: Print in required format
	solution = board.String()

	return solution, err
}

func FirstSolution(size uint8, tiles []Tile) (solution string, err error) {
	solutions, err := Solve(size, tiles)

	if err != nil {
		return solution, err
	}

	board := <-solutions
	solution, err = BoardToSolution(board, tiles)
	return solution, err
}