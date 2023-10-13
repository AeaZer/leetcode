package dp

import "math"

// 120. 三角形最小路径和
// https://leetcode.cn/problems/triangle/description/

/*给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与
上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。*/

/*示例 1：
输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

示例 2：
输入：triangle = [[-10]]
输出：-10*/

/*f(0) = triangle[0][0]
f(0,0) = f(0)+triangle[1][0], f(0,1) = f(0)+triangle[1][1]
f(0,0,0) = f(0,0)+ triangle[2][0], f(0,0,1) = f(0,0)+ triangle[2][1], f(0,1,1) = f(0,0)+ triangle[2][1]
f(0,0,0,0) = f(0, 0, 0)+ triangle[3][0], f(0,0,0,1) = f(0,0)+ triangle[2][1], f(0,1,1) = f(0,0)+ triangle[2][1]*/
// 动态规划转移方程：
func minimumTotal(triangle [][]int) int {
	tl := len(triangle)
	dp := make([][]int, tl)
	for i := 0; i < tl; i++ {
		dp[i] = make([]int, tl)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < tl; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for i := 0; i < tl; i++ {
		ans = min(ans, dp[tl-1][i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
