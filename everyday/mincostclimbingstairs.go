package everyday

// 746. 使用最小花费爬楼梯
// https://leetcode.cn/problems/min-cost-climbing-stairs/description/?envType=daily-question&envId=2023-12-17

func minCostClimbingStairs(cost []int) int {
	cl := len(cost)
	if cl == 0 {
		return 0
	}
	dp := make([]int, cl+1)
	for i := 2; i <= cl; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])		
	}
	return dp[cl]
}