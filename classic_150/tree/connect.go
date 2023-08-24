package tree

// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/

/*填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
初始状态下，所有 next 指针都被设置为 NULL 。*/

type connectNode struct {
	Val   int
	Left  *connectNode
	Right *connectNode
	Next  *connectNode
}

/**
 * Definition for a connectNode.
 * type connectNode struct {
 *     Val int
 *     Left *connectNode
 *     Right *connectNode
 *     Next *connectNode
 * }
 */

func connect(root *connectNode) *connectNode {
	deepStore := make(map[int][]*connectNode)
	var recursive func(root *connectNode, depth int)
	recursive = func(root *connectNode, depth int) {
		if root == nil {
			return
		}
		depth++
		nodes, ok := deepStore[depth]
		if ok {
			// 当一行有多个的时候当然是左边的下一个节点总是数组中的最后一个节点
			root.Next = nodes[len(nodes)-1]
		}
		deepStore[depth] = append(nodes, root)
		// 左边的 next 是右边的 node 那么我们需要先递归右节点，将右节点 append 进 deepStore[deep] 中
		recursive(root.Right, depth)
		recursive(root.Left, depth)
	}
	recursive(root, 0)
	return root
}
