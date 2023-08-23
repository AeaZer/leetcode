package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	assert := assert.New(t)

	l1 := newListNode([]int{1, 2, 3, 4, 5})
	node := removeNthFromEnd(l1, 2)
	assert.Equal([]int{1, 2, 3, 5}, rangeListNode(*node))

	l1 = newListNode([]int{1})
	node = removeNthFromEnd(l1, 1)
	assert.Nil(node)

	l1 = newListNode([]int{1, 2})
	node = removeNthFromEnd(l1, 1)
	assert.Equal([]int{1}, rangeListNode(*node))
}
