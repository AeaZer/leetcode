package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanFinish(t *testing.T) {
	assert := assert.New(t)

	finish := canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}})
	assert.True(finish)

	finish = canFinish(3, [][]int{{1, 4}})
	assert.True(finish)

	finish = canFinish(7, [][]int{{1, 0}, {0, 3}, {0, 2}, {3, 2}, {2, 5}, {4, 5}, {5, 6}, {2, 4}})
	assert.True(finish)
}
