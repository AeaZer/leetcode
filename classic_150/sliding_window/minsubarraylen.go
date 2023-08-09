package sliding_window

import "math"

/*给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的
连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。*/

// https://leetcode.cn/problems/minimum-size-subarray-sum/

// 滑動窗口
func minSubArrayLen(target int, nums []int) int {
	nl := len(nums)
	if nl == 0 {
		return 0
	}

	minLen := math.MaxInt
	lp, rp := 0, 0
	sum := 0
	for rp < nl {
		sum += nums[rp]
		for sum >= target {
			minLen = min(minLen, rp-lp+1)
			sum -= nums[lp]
			lp++
		}
		rp++
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
