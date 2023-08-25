package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	depthMap := make(map[int][]int)
	var dfs func(rNode *TreeNode, depth int)
	dfs = func(rNode *TreeNode, depth int) {
		if rNode == nil {
			return
		}
		depthMap[depth] = append(depthMap[depth], rNode.Val)
		dfs(rNode.Left, depth+1)
		dfs(rNode.Right, depth+1)
	}
	dfs(root, 0)
	ret := make([][]int, 0, len(depthMap))
	for i := 0; ; i++ {
		values, ok := depthMap[i]
		if !ok {
			break
		}
		ret = append(ret, values)
	}
	return ret
}
