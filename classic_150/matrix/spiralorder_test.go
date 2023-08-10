package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	res := spiralOrder(matrix)
	assert.Equal([]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}, res)

	matrix = [][]int{{6, 9, 7}}
	res = spiralOrder(matrix)
	assert.Equal([]int{6, 9, 7}, res)

	matrix = [][]int{{6}, {9}, {7}}
	res = spiralOrder(matrix)
	assert.Equal([]int{6, 9, 7}, res)
}
