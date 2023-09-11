package tree

// 222. 完全二叉树的节点个数
// https://leetcode.cn/problems/count-complete-tree-nodes/

/*给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。
若最底层为第 h 层，则该层包含 1~ 2h 个节点。*/

// 进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？

// 那我们就进阶一下，单纯遍历太无聊了，根据完全二叉树的特性来做
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dummyRoot := root
	maxD := depth(dummyRoot, 0)
	var (
		count     int
		nullRight bool
	)
	var helper func(r *TreeNode, d int)
	helper = func(r *TreeNode, d int) {
		if r == nil {
			if d == maxD {
				nullRight = true
			}
			return
		}
		if d == maxD {
			count++
		}
		if !nullRight {
			helper(r.Left, d+1)
			helper(r.Right, d+1)
		}
	}
	helper(root, 1)
	return 1<<(maxD-1) + count - 1
}

func depth(root *TreeNode, d int) int {
	if root == nil {
		return d
	}
	return depth(root.Left, d+1)
}
