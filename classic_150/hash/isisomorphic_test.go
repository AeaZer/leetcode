package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	assert := assert.New(t)

	ok := isIsomorphic("egg", "add")
	assert.True(ok)

	ok = isIsomorphic("foo", "bar")
	assert.False(ok)
}
