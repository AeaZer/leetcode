package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	assert := assert.New(t)

	finish := canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}})
	assert.True(finish)

	finish = canFinish(3, [][]int{{1, 4}})
	assert.True(finish)

	finish = canFinish2(7, [][]int{{1, 3}, {2, 3}})
	assert.True(finish)
}
