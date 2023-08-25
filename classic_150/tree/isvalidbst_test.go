package tree

import "testing"

func TestIsValidBST(t *testing.T) {
	node := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	validBST := isValidBST(node)
	t.Log(validBST)
}
