package tree

import "math"

// 98. 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/

/*给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。*/

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	val := node.Val
	if val <= lower || val >= upper {
		return false
	}
	return helper(node.Right, val, upper) && helper(node.Left, lower, val)
}
