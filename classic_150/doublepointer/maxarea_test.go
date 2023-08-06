package doublepointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {
	assert := assert.New(t)

	area := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	assert.Equal(49, area)
}
