package tree

// 129. 求根节点到叶节点数字之和
// https://leetcode.cn/problems/sum-root-to-leaf-numbers/

/*给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func arrToNum(arr []int) int {
	var res int
	for i := range arr {
		res = res*10 + arr[i]
	}
	return res
}

func sumNumbers(root *TreeNode) int {
	var res int
	var recursive func(root *TreeNode, values []int)
	recursive = func(root *TreeNode, values []int) {
		if root == nil {
			return
		}
		values = append(values, root.Val)
		// 叶子节点
		if root.Right == nil && root.Left == nil {
			res += arrToNum(values)
		}
		// 深度遍历
		recursive(root.Left, values)
		recursive(root.Right, values)
	}
	recursive(root, []int{})
	return res
}
