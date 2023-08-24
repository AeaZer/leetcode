package tree

// 112. 路径总和
// https://leetcode.cn/problems/path-sum/

// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
// 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
// 叶子节点 是指没有子节点的节点。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	var found bool
	var recursive func(root *TreeNode, targetSum int)
	recursive = func(root *TreeNode, targetSum int) {
		if found || root == nil {
			return
		}
		// root.Val- targetSum == 0 且是叶子节点
		if root.Val == targetSum && root.Right == nil && root.Left == nil {
			found = true
			return
		}
		// 深度遍历
		recursive(root.Left, targetSum-root.Val)
		recursive(root.Right, targetSum-root.Val)
	}
	recursive(root, targetSum)
	return found
}
