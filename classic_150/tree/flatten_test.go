package tree

import (
	"testing"
)

func TestFlatten(t *testing.T) {

	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(node)

	node = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	flatten(node)
	t.Log(11)
}
