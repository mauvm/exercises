package puzzle

import "testing"
import "github.com/stretchr/testify/assert"
import "strconv"
import "strings"
import "io/ioutil"

func readFile(t *testing.T, file string) string {
	data, err := ioutil.ReadFile(file)
	assert.Equal(t, nil, err)
	return string(data)
}

func readInputFile(t *testing.T, file string) (uint8, []string) {
	data := readFile(t, file)
	lines := strings.Split(data, "\n")
	if len(lines) == 0 {
		return 0, []string{}
	}
	amount, _ := strconv.Atoi(lines[0])
	return uint8(amount), lines[1:]
}

func TestFirstSolution(t *testing.T) {
	size, lines := readInputFile(t, "../fixtures/input_2.txt")
	tiles, err := LinesToTiles(lines)
	assert.Equal(t, nil, err)
	_, err = FirstSolution(size, tiles)
	assert.Equal(t, nil, err)
}
