package tree

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

func invertTree(root *TreeNode) *TreeNode {
	dummyRoot := root
	var recursive func(root *TreeNode)
	recursive = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		recursive(root.Left)
		recursive(root.Right)
	}
	recursive(dummyRoot)
	return root
}
