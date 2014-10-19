package main

import "github.com/mauvm/exercises/dailyprogrammer/183_intermediate_edge_matching/puzzle"
import "strconv"
import "strings"
import "io/ioutil"
import "fmt"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(file string) string {
	data, err := ioutil.ReadFile(file)
	checkError(err)
	return string(data)
}

func ReadInputFile(file string) (uint8, []string) {
	data := readFile(file)
	lines := strings.Split(data, "\n")
	if len(lines) == 0 {
		return 0, []string{}
	}
	amount, _ := strconv.Atoi(lines[0])
	return uint8(amount), lines[1:]
}

func LinesToTiles(lines []string) []puzzle.Tile {
	tiles := []puzzle.Tile{}
	for _, line := range lines {
		tile, err := puzzle.TileFromString(line)
		checkError(err)
		tiles = append(tiles, tile)
	}
	return tiles
}

func main() {
	size, lines := ReadInputFile("fixtures/input_1.txt")
	tiles := LinesToTiles(lines)
	solution, err := puzzle.FirstSolution(size, tiles)
	checkError(err)
	fmt.Println(solution)
}
