package tree

// 230. 二叉搜索树中第K小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode)
	var res int
	dfs = func(root *TreeNode) {
		if root == nil || k == 0 {
			return
		}
		dfs(root.Left)
		if k == 1 {
			res = root.Val
		}
		k--
		dfs(root.Right)
	}
	dfs(root)
	return res
}
