package binarysearch

// 35. 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/description/?envType=study-plan-v2&envId=top-interview-150

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if mid == left {
			break
		}
		if nums[mid] > target {
			right = mid
			continue
		}
		left = mid
	}
	if nums[right] < target {
		return right + 1
	}
	if nums[left] < target && nums[right] >= target {
		return right
	}
	if left == 0 {
		return 0
	}
	return left - 1
}

func searchInsert2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
			continue
		}
		right = mid - 1
	}
	return left
}
