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

import "strconv"
import "regexp"
import "strings"
import "io/ioutil"
import "sort"
import "fmt"

type SemanticVersion struct {
	major, minor, patch int
	label, metadata string
}

func (v SemanticVersion) ToString() string {
	version := strconv.Itoa(v.major)
	version += "." + strconv.Itoa(v.minor)
	version += "." + strconv.Itoa(v.patch)
	if v.label != "" {
		version += "-" + v.label
	}
	if v.metadata != "" {
		version += "+" + v.metadata
	}
	return version
}

type SemanticVersions []SemanticVersion

func (slice SemanticVersions) Len() int {
	return len(slice)
}

func (slice SemanticVersions) Less(i, j int) bool {
	v1 := slice[i]
	v2 := slice[j]
	if v1.major != v2.major {
		return v1.major < v2.major
	}
	if v1.minor != v2.minor {
		return v1.minor < v2.minor
	}
	if v1.patch != v2.patch {
		return v1.patch < v2.patch
	}
	if len(v1.label) > 0 || len(v2.label) > 0 {
		return v1.label > v2.label
	}
	if len(v1.metadata) > 0 || len(v2.metadata) > 0 {
		return v1.metadata > v2.metadata
	}
	return true // Equal
}

func (slice SemanticVersions) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func SemanticVersionFromString(input string) SemanticVersion {
	re := regexp.MustCompile("^([0-9]+)\\.([0-9]+)\\.([0-9]+)(-[^\\+]+)?(\\+.+)?$")
	matches := re.FindAllStringSubmatch(input, -1)[0]
	version := SemanticVersion{}

	version.major, _ = strconv.Atoi(matches[1])
	version.minor, _ = strconv.Atoi(matches[2])
	version.patch, _ = strconv.Atoi(matches[3])

	if len(matches[4]) > 0 {
		version.label = matches[4][1:] // Strip off '-'
	}
	if len(matches[5]) > 0 {
		version.metadata = matches[5][1:] // Strip off '+'
	}

	return version
}

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

func LinesToVersions(lines []string) SemanticVersions {
	versions := SemanticVersions{}
	for _, line := range lines {
		version := SemanticVersionFromString(line)
		versions = append(versions, version)
	}
	return versions
}

func SortVersions(versions SemanticVersions) string {
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
