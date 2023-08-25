package tree

// 103. 二叉树的锯齿形层序遍历
// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/

/*给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。
（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	depthMap := make(map[int][]int)
	var dfs func(rNode *TreeNode, depth int)
	dfs = func(rNode *TreeNode, depth int) {
		if rNode == nil {
			return
		}
		if depth%2 == 0 {
			depthMap[depth] = append(depthMap[depth], rNode.Val)
		} else {
			depthMap[depth] = append([]int{rNode.Val}, depthMap[depth]...)
		}
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
