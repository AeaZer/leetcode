package tree

import "testing"

func TestMinDepth(t *testing.T) {
	node := &TreeNode{
		Val: 236,
		Left: &TreeNode{
			Val: 104,
			Right: &TreeNode{
				Val: 227,
			},
		},
		Right: &TreeNode{
			Val: 701,
			Right: &TreeNode{
				Val: 911,
			},
		},
	}

	i := minDepth(node)
	t.Log(i)
}
