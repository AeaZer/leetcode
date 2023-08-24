package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	deepStore := make(map[int][]*Node)
	var recursive func(root *Node, depth int)
	recursive = func(root *Node, depth int) {
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
