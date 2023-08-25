package tree

// 199. 二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/

// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	depthMap := make(map[int]bool)
	ret := make([]int, 0)
	var dfs func(rNode *TreeNode, depth int)
	dfs = func(rNode *TreeNode, depth int) {
		if rNode == nil {
			return
		}
		if !depthMap[depth] {
			ret = append(ret, rNode.Val)
			depthMap[depth] = true
		}
		dfs(rNode.Right, depth+1)
		dfs(rNode.Left, depth+1)
	}
	dfs(root, 1)
	return ret
}
