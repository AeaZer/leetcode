package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	assert := assert.New(t)

	l1 := newListNode([]int{1, 4, 3, 2, 5, 2})
	node := partition(l1, 3)
	assert.Equal([]int{1, 2, 2, 4, 3, 5}, rangeListNode(*node))
}
