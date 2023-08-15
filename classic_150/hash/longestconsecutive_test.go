package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	assert := assert.New(t)

	maxLength := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	assert.Equal(4, maxLength)

	maxLength = longestConsecutive([]int{1, 2, 0, 1})
	assert.Equal(3, maxLength)
}
