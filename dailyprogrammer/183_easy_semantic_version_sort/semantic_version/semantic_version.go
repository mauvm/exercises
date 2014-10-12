package semantic_version

import "strconv"
import "regexp"

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

// Implement sortInterface
// Source: http://nerdyworm.com/blog/2013/05/15/sorting-a-slice-of-structs-in-go/

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
	return false // Equal
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
