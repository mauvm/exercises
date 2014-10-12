// Assignment: http://www.reddit.com/r/dailyprogrammer/comments/2igfj9/10062014_challenge_183_easy_semantic_version_sort/

// Learned about:
// - Unit testing
// - Structures
// - Slices and manipulation
// - Sort interface
// - Regular expressions
// - File reading
// - Ranges
// - String manipulation
// - Sorting

package main

import "github.com/mauvm/exercises/dailyprogrammer/183_easy_semantic_version_sort/semantic_version"
import "strconv"
import "strings"
import "io/ioutil"
import "sort"
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

func ReadVersionFile(file string) []string {
	data := readFile(file)
	lines := strings.Split(data, "\n")
	if len(lines) == 0 {
		return []string{}
	}
	amount, _ := strconv.Atoi(lines[0])
	if len(lines) < amount + 1 {
		return lines[1:]
	}
	return lines[1:amount+1]
}

func LinesToVersions(lines []string) semantic_version.SemanticVersions {
	versions := semantic_version.SemanticVersions{}
	for _, line := range lines {
		version := semantic_version.SemanticVersionFromString(line)
		versions = append(versions, version)
	}
	return versions
}

func SortVersions(versions semantic_version.SemanticVersions) string {
	result := ""
	if len(versions) == 0 {
		return result
	}
	sort.Sort(versions)
	for _, version := range versions {
		result += version.ToString() + "\n"
	}
	result = result[:len(result)-1] // Remove trailing newline
	return result
}

func main() {
	lines := ReadVersionFile("fixtures/input_1.txt")
	versions := LinesToVersions(lines)
	output := SortVersions(versions)
	fmt.Print(output)
}
