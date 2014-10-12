package semantic_version

import "testing"
import "github.com/stretchr/testify/assert"

func TestSemanticVersionToString(t *testing.T) {
	v := SemanticVersion{1, 2, 3, "alpha", "20141010"}
	assert.Equal(t, v.ToString(), "1.2.3-alpha+20141010")
}

func TestSemanticVersionFromString(t *testing.T) {
	v := SemanticVersionFromString("1.2.3-alpha+20141010")
	assert.Equal(t, v.major, 1)
	assert.Equal(t, v.minor, 2)
	assert.Equal(t, v.patch, 3)
	assert.Equal(t, v.label, "alpha")
	assert.Equal(t, v.metadata, "20141010")

	v = SemanticVersionFromString("1000.0.11+20111111")
	assert.Equal(t, v.major, 1000)
	assert.Equal(t, v.minor, 0)
	assert.Equal(t, v.patch, 11)
	assert.Equal(t, v.label, "")
	assert.Equal(t, v.metadata, "20111111")
}
