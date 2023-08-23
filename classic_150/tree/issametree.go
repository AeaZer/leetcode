package tree

// 100. 相同的树
// https://leetcode.cn/problems/same-tree/

/*给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
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
		recursive(p.Right, q.Right)
		recursive(p.Left, q.Left)
	}
	recursive(p, q)
	return same
}
