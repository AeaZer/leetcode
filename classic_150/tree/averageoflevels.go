package tree

// 637. 二叉树的层平均值
// https://leetcode.cn/problems/average-of-levels-in-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	store := make(map[int][2]int)
	var dfs func(rNode *TreeNode, depth int)
	dfs = func(rNode *TreeNode, depth int) {
		if rNode == nil {
			return
		}
		values := store[depth]
		values[0]++
		values[1] += rNode.Val
		store[depth] = values
		dfs(rNode.Left, depth+1)
		dfs(rNode.Right, depth+1)
	}
	dfs(root, 0)
	ret := make([]float64, 0, len(store))
	for i := 0; ; i++ {
		values, ok := store[i]
		if !ok {
			break
		}
		ret = append(ret, float64(values[1])/float64(values[0]))
	}
	return ret
}
