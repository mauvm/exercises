package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestVersionSort(t *testing.T) {
	lines := ReadVersionFile("fixtures/input_1.txt")
	versions := LinesToVersions(lines)
	output := SortVersions(versions)
	expectedOutput := readFile("fixtures/output_1.txt")
	assert.Equal(t, output, expectedOutput)
}
