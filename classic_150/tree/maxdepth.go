package tree

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

/*给定一个二叉树 root ，返回其最大深度。
二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var ret int
	var recursive func(root *TreeNode, depth int)
	recursive = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		depth++
		if depth > ret {
			ret = depth
		}
		recursive(root.Left, depth)
		recursive(root.Right, depth)
	}
	recursive(root, 0)
	return ret
}

func minDepth(root *TreeNode) int {
	var ret int
	var recursive func(root *TreeNode, depth int)
	recursive = func(root *TreeNode, depth int) {
		if root == nil || (ret != 0 && depth > ret) {
			return
		}
		if root.Left == nil && root.Right == nil {
			ret = depth
		}
		recursive(root.Left, depth+1)
		recursive(root.Right, depth+1)
	}
	recursive(root, 1)
	return ret
}
