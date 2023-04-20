package main

type TreeNode struct {
	left, right *TreeNode
	value       int
}

func newTree() *TreeNode {
	return &TreeNode{
		value: 3,
		left: &TreeNode{
			value: 1,
			right: &TreeNode{
				value: 2,
			},
		},
		right: &TreeNode{
			value: 4,
		},
	}
}

var res []int

func traverseMiddle(treeNode *TreeNode) []int {
	if treeNode == nil {
		return res
	}
	traverseMiddle(treeNode.left)
	res = append(res, treeNode.value)
	traverseMiddle(treeNode.right)
	return res
}

func (t *TreeNode) searchKth(k int) (value int, ok bool) {
	values := traverseMiddle(t)
	if len(values) < k {
		return
	}
	return values[k-1], true
}
