package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)

	its := twoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal([]int{0, 1}, its)
}
