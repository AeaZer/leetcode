package tree

// 101. 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/

// 给你一个二叉树的根节点 root ，检查它是否轴对称。

func isSymmetric(root *TreeNode) bool {
	same := true
	var recursive func(p *TreeNode, q *TreeNode)
	recursive = func(p *TreeNode, q *TreeNode) {
		if !same || (p == nil && q == nil) {
			return
		}
		if (p == nil && q != nil) || (p != nil && q == nil) || (p != nil && q != nil && p.Val != q.Val) {
			same = false
			return
		}
		// 我们发现中心轴对成，除了根节点之后，左子树和右子树相同，右子树和左子树相同
		recursive(p.Right, q.Left)
		recursive(p.Left, q.Right)
	}
	recursive(root.Left, root.Right)
	return same
}
