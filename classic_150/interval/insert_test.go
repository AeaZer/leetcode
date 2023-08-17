package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert := assert.New(t)

	res := insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	assert.Equal([][]int{{1, 5}, {6, 9}}, res)
}
