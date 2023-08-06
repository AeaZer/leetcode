package doublepointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)

	indexes := twoSum([]int{2, 3, 4}, 6)
	assert.Equal([]int{1, 3}, indexes)

	indexes = twoSum([]int{-1, 0}, -1)
	assert.Equal([]int{1, 2}, indexes)

	indexes = twoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal([]int{1, 2}, indexes)
}
