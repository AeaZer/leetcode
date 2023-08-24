package tree

// 114. 二叉树展开为链表
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/

/*给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	var recursive func(root *TreeNode)
	recursive = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Right != nil {
			right := root.Right
			root.Right = root.Left
			dummyCurrentRoot := root
			// 移动过来的左子树可能有很多，所以需要找到最下最右边不为 nil 的节点，将 node.Right 指向原来的有右子树
			for dummyCurrentRoot.Right != nil {
				dummyCurrentRoot = dummyCurrentRoot.Right
			}
			dummyCurrentRoot.Right = right
		} else {
			root.Right = root.Left
		}
		// 最后别忘记清空左节点
		root.Left = nil
		// 进入下一个迭代
		recursive(root.Right)
	}
	dummyRoot := root
	recursive(dummyRoot)
}
