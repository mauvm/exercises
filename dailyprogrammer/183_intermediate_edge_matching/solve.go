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

func readInputFile(file string) (uint8, []string) {
	data := readFile(file)
	lines := strings.Split(data, "\n")
	if len(lines) == 0 {
		return 0, []string{}
	}
	amount, _ := strconv.Atoi(lines[0])
	return uint8(amount), lines[1:]
}

func main() {
	size, lines := readInputFile("fixtures/input_1.txt")
	tiles, err := puzzle.LinesToTiles(lines)
	checkError(err)
	solution, err := puzzle.FirstSolution(size, tiles)
	checkError(err)
	fmt.Println(solution)
}
