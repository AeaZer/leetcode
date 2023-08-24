package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	assert := assert.New(t)

	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	res := sumNumbers(node)
	assert.Equal(25, res)
}
