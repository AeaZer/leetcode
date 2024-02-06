package dp

/*你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。*/

// 1 <= nums.length <= 100

func rob(nums []int) int {
	nl := len(nums)
	dp := make([]int, nl)
	if nl == 1 {
		return nums[0]
	}
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < nl; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[nl-1]
}

func rob2(nums []int) int {
	nl := len(nums)
	if nl == 1 {
		return nums[0]
	}
	g1, g2 := nums[0], max(nums[0], nums[1])
	for i := 2; i < nl; i++ {
		g1, g2 = g2, max(g1+nums[i], g2)
	}
	return g2
}
