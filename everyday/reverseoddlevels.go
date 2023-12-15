package everyday

// 2415. 反转二叉树的奇数层
// https://leetcode.cn/problems/reverse-odd-levels-of-binary-tree/description

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func reverseOddLevels(root *TreeNode) *TreeNode {
	var dfs func (*TreeNode, int)
	oddDepthTreeMap := make(map[int][]*TreeNode, 0)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth % 2 != 0 {
			oddDepthTreeMap[depth] = append(oddDepthTreeMap[depth], node)
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	for _, nodes := range oddDepthTreeMap {
		nl := len(nodes)
		for i := 0; i < nl/2; i++ {
			nodes[i].Val, nodes[nl-i-1].Val = nodes[nl-i-1].Val, nodes[i].Val
		}
	}
	return root
 }

