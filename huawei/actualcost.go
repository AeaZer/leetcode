package huawei

import (
	"strconv"
)

// 题目：计程车打表真实价格
// 场景设计是程序员小 i 打车，发现计程车的计时器有问题，开车的师傅说是不喜欢数字 4，所有含有数字4的数字都会跳开，求实际的花费

func actualCost(cost int64) int64 {
	costStr := strconv.FormatInt(cost, 10)
	y := len(costStr)
	dp := make([]int64, y)
	dp[0] = 0
	for i := 1; i < y; i++ {
		dp[i] = 10*dp[i-1] + 1
	}
	var fourC int64
	for i := y - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(costStr[i]))
		fourC += dp[y-i-1] * int64(num)
		if num > 4 {
			fourC++
		}
	}
	return cost - fourC
}
