package tree

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

/*给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历，
inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。*/

func findIndex(arr []int, num int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}

func buildTreeWithPreorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	index := findIndex(inorder, preorder[0])
	root.Left = buildTreeWithPreorder(preorder[1:index+1], inorder[:index])
	root.Right = buildTreeWithPreorder(preorder[index+1:], inorder[index+1:])
	return root
}

func buildTreePostorder(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	lastE := postorder[len(postorder)-1]
	root := &TreeNode{Val: lastE}
	index := findIndex(inorder, lastE)
	root.Left = buildTreePostorder(inorder[:index], postorder[:index])
	root.Right = buildTreePostorder(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
