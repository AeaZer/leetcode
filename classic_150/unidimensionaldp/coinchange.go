package unidimensionaldp

import "math"

// 322. 零钱兑换
// https://leetcode.cn/problems/coin-change/description/?envType=study-plan-v2&envId=top-interview-150

/*给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。
*/

// 一维动态规划转移方程 dp[money] = min(dp[money-coin[0]], dp[money-coin[1]], ..., dp[money-coin[i]]) + 1

// coins = [1, 2, 5]
// f(0) = 0
// f(1) = min(f(1-1)+1, f(1-2)+1, f(1-5)+1)
// f(2) = min(f(2-1), f(2-2)+1, f(2-5)+1)
// f(3) = min(f(3-1), f(3-2)+1, f(3-5)+1)

// coins = [3, 5]
// f(0) = 0
// f(1) = min(f(1-3) + f(1-5))+1
// f(2) = min(f(2-3) + f(2-5))+1
// f(3) = min(f(3-3) + f(3-5))+1 = 1
// f(6) = min(f(6-3) + f(6-5))+1

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for _, coin := range coins {
			if i-coin < 0 || dp[i-coin] == math.MaxInt {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
