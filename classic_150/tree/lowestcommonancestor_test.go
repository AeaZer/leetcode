package tree

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	p, q := &TreeNode{Val: 9}, &TreeNode{Val: 20}
	node := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val:   15,
			Left:  p,
			Right: q,
		},
	}
	treeNode := lowestCommonAncestor(node, q, p)
	t.Log(treeNode.Val)
}
