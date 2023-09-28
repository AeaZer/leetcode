package binarysearch

// 34. 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	i := -1
	for left <= right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			i = mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// not search successful
	res := []int{i, i}
	if i == -1 {
		return res
	}
	for lower := i - 1; lower >= 0; lower-- {
		if nums[lower] != target {
			break
		}
		res[0] = lower
	}
	for higher := i + 1; higher < len(nums); higher++ {
		if nums[higher] != target {
			break
		}
		res[1] = higher
	}
	return res
}
