package string_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert := assert.New(t)

	strings := []string{"", ""}
	commonPrefix := longestCommonPrefix(strings)
	assert.Equal("", commonPrefix)

	strings = []string{"flower", "flow", "flight"}
	commonPrefix = longestCommonPrefix(strings)
	assert.Equal("fl", commonPrefix)

	strings = []string{"dog", "racecar", "car"}
	commonPrefix = longestCommonPrefix(strings)
	assert.Equal("", commonPrefix)
}
