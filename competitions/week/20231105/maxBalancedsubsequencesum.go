package _0231105

/*输入：nums = [3,3,5,6]
输出：14
解释：这个例子中，选择子序列 [3,5,6] ，下标为 0 ，2 和 3 的元素被选中。
nums[2] - nums[0] >= 2 - 0 。
nums[3] - nums[2] >= 3 - 2 。
所以，这是一个平衡子序列，且它的和是所有平衡子序列里最大的。
包含下标 1 ，2 和 3 的子序列也是一个平衡的子序列。
最大平衡子序列和为 14 。*/

func maxBalancedSubsequenceSum(nums []int) int64 {
	numsLength := len(nums)
	dp := make([]int, numsLength)
	dp[0] = nums[0]
	for i := 1; i < numsLength; i++ {
		dpValue := dp[i-1]
		for j := 0; j < i; j++ {
			if nums[i]-nums[j] >= i-j && nums[i]+dp[j] > dpValue {
				dpValue = nums[i] + dp[j]
			}
		}
		dp[i] = max(dpValue, nums[i])
	}
	return int64(dp[numsLength-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
