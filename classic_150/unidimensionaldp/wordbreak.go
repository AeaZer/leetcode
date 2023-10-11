package unidimensionaldp

// 139. 单词拆分
// https://leetcode.cn/problems/word-break/description/?envType=study-plan-v2&envId=top-interview-150

// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

// 动态规划转移方程 dp[i] = dp[j] && check(s[j..i−1])
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		m[wordDict[i]] = true
	}
	sl := len(s)
	dp := make([]bool, sl+1)
	dp[0] = true
	for i := 1; i <= sl; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[sl]
}
