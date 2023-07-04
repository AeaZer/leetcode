package tree

import "sort"

type node struct {
	left, right *node
	value       int
}

func (tree *node) ValueInBST(value int) bool {
	if tree == nil {
		return false
	}
	if value < tree.value {
		return tree.left.ValueInBST(value)
	}
	if value > tree.value {
		return tree.right.ValueInBST(value)
	}
	return value == tree.value
}

func (tree *node) InsertValue(value int) *node {
	if tree == nil {
		return &node{value: value}
	}
	if value <= tree.value {
		tree.left = tree.left.InsertValue(value)
	}
	if value > tree.value {
		tree.right = tree.right.InsertValue(value)
	}
	return tree
}

// SearchKValue 获取第 K 大小的值
func (tree *node) SearchKValue(k int) (int, bool) {
	if k <= 0 {
		return 0, false
	}
	slice := RangeBSTToSlice(tree)
	if len(slice) < k {
		return 0, false
	}
	return slice[k-1], true
}

func (tree *node) beforeTraverse(res []int) []int {
	if tree == nil {
		return res
	}
	res = append(res, tree.value)
	tree.left.beforeTraverse(res)
	tree.right.beforeTraverse(res)
	return res
}

func (tree *node) BeforeTraverse() []int {
	return tree.beforeTraverse([]int{})
}

func (tree *node) afterTraverse(res []int) []int {
	if tree == nil {
		return res
	}
	tree.left.afterTraverse(res)
	tree.right.afterTraverse(res)
	res = append(res, tree.value)
	return res
}

func (tree *node) AfterTraverse() []int {
	return tree.afterTraverse([]int{})
}

func newBSTWithSlice(nums []int, needSort bool) *node {
	if len(nums) == 0 {
		return nil
	}
	if needSort {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
	}
	midIndex := len(nums) / 2
	root := &node{value: nums[midIndex]}
	root.left = newBSTWithSlice(nums[:midIndex], false)
	root.right = newBSTWithSlice(nums[midIndex+1:], false)
	return root
}

func NewBSTWithSlice(nums []int) *node {
	return newBSTWithSlice(nums, true)
}

func rangeBSTToSlice(tree *node, res []int) []int {
	if tree == nil {
		return res
	}
	res = rangeBSTToSlice(tree.left, res)
	res = append(res, tree.value)
	res = rangeBSTToSlice(tree.right, res)
	return res
}

func RangeBSTToSlice(tree *node) []int {
	return rangeBSTToSlice(tree, []int{})
}
