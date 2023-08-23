package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	assert := assert.New(t)

	l1 := newListNode([]int{1, 2, 3, 3, 4, 4, 5})
	node := deleteDuplicates(l1)
	assert.Equal([]int{1, 2, 5}, rangeListNode(*node))
}
