package string_array

import "sort"

/*给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。*/

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

// 思路： 贪心 + 动态算出第 N 天卖出的最大利润，然后取最大值
func maxProfit2(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}

	dp := make([]int, length)
	dp[0] = 0
	for i := 1; i < length; i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 || dp[i-1]+profit > 0 {
			dp[i] = dp[i-1] + profit
			continue
		}
		dp[i] = 0
	}

	sort.Ints(dp)
	return dp[length-1]
}

func maxProfit1(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}

	var (
		beforeProfit     int
		currentMaxProfit int
	)
	for i := 1; i < length; i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 || beforeProfit+profit > 0 {
			beforeProfit += profit
			if beforeProfit > currentMaxProfit {
				currentMaxProfit = beforeProfit
			}
			continue
		}
		beforeProfit = 0
	}

	return currentMaxProfit
}

/*给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
返回 你能获得的 最大 利润 。*/

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

// 究极贪心，亏就卖
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}

	var (
		currentProfit int
		totalProfit   int
	)
	for i := 1; i < length; i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			currentProfit += profit
			if i == length-1 {
				totalProfit += currentProfit
			}
			continue
		}
		totalProfit += currentProfit
		currentProfit = 0
	}
	return totalProfit
}
