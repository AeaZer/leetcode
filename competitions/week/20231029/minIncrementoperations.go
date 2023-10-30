package _0231029

func minIncrementOperations(nums []int, k int) int64 {
	n := len(nums)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}

	minOps := int64(^uint(0) >> 1)
	for i := 0; i+2 < n; i++ {
		if nums[i+2] < k {
			continue
		}
		ops := dp[i+2] - dp[i] - 2*(i+1) + (n - i - 2) - (nums[i+2] - k)
		if int64(ops) < minOps {
			minOps = int64(ops)
		}
	}
	return minOps
}
