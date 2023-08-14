package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	assert := assert.New(t)

	ok := isAnagram("anagram", "nagaram")
	assert.True(ok)

	ok = isAnagram("rat", "car")
	assert.False(ok)
}
