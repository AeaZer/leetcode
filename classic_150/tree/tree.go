package tree

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newBSTWithSlice(nums []int, needSort bool) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if needSort {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
	}
	midIndex := len(nums) / 2
	root := &TreeNode{Val: nums[midIndex]}
	root.Left = newBSTWithSlice(nums[:midIndex], false)
	root.Right = newBSTWithSlice(nums[midIndex+1:], false)
	return root
}

func NewBSTWithSlice(nums []int) *TreeNode {
	return newBSTWithSlice(nums, true)
}
