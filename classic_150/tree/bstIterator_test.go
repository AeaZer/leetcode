package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBSTIterator(t *testing.T) {
	assert := assert.New(t)

	node := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	iterator := Constructor(node)
	next := iterator.Next()
	assert.Equal(3, next)
}
