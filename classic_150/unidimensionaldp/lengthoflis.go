package unidimensionaldp

// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/description/?envType=study-plan-v2&envId=top-interview-150

/*给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。*/

func lengthOfLIS(nums []int) int {
	nl := len(nums)
	if nl == 0 {
		return 0
	}
	dp := make([]int, nl)
	dp[0] = 1
	var maxIncrease = 1
	for i := 1; i < nl; i++ {
		if nums[i]-nums[i-1] > 0 {
			dp[i] = dp[i-1] + 1
			if dp[i] > maxIncrease {
				maxIncrease = dp[i]
			}
			continue
		}
		dp[i] = 1
	}
	return maxIncrease
}
