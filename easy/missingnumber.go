package main

import (
	"sort"
)

// 268. 丢失的数字
// https://leetcode.cn/problems/missing-number/description/

// 给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

func missingNumber(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if mid == nums[mid] {
			left = mid + 1
			continue
		}
		right = mid - 1
	}
	return left
}
