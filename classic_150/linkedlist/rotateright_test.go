package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	assert := assert.New(t)

	l1 := newListNode([]int{1, 2, 3})
	node := rotateRight(l1, 1)
	assert.Equal([]int{3, 1, 2}, rangeListNode(*node))
}
