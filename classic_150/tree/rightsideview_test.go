package tree

import "testing"

func TestRightSideView(t *testing.T) {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	view := rightSideView(node)
	t.Log(view)
}
