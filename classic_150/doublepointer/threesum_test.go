package doublepointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	assert := assert.New(t)

	res := threeSum([]int{-4, -2, -1})
	assert.Empty(res)

	res = threeSum([]int{-2, 0, 1, 1, 2})
	assert.Equal([][]int{{-2, 0, 2}, {-2, 1, 1}}, res)

	res = threeSum([]int{-1, 0, 1, 2, -1, -4})
	assert.Equal([][]int{{-1, -1, 2}, {-1, 0, 1}}, res)

	res = threeSum([]int{0, 1, 1})
	assert.Empty(res)

	res = threeSum([]int{0, 0, 0})
	assert.Equal([][]int{{0, 0, 0}}, res)

	res = threeSum([]int{-1, 0, 1, 0})
	assert.Equal([][]int{{-1, 0, 1}}, res)
}
