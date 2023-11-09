package dp

import "math"

// 72. 编辑距离
// https://leetcode.cn/problems/edit-distance/description/?envType=study-plan-v2&envId=top-interview-150

/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
*/

/*输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')*/
// 思考：已知这个是一道动态规划的题目，那么我觉得可以这样想：
// word1[n]  => word2[0] 1 2 2 3 4
// word1[n]  => word2[1] word1[n][0] + 1, word1[n][1] + 0, word1[n][2] + 1, word1[n][3] + 1, word1[n][4] + 1
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		if word1[j] != word2[0] {
			dp[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[j] != word2[i] {
				dp[i][j] = dp[i-1][j] + 1
				continue
			}
			dp[i][j] = dp[i-1][j]
		}
	}
	ans := math.MaxInt
	for j := 0; j < n; j++ {
		if dp[m-1][j] < ans {
			ans = dp[m-1][j]
		}
	}
	return ans
}
