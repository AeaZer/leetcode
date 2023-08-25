package tree

import "testing"

func TestGetMinimumDifference(t *testing.T) {
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
	res := getMinimumDifference(node)
	t.Log(res)
}
