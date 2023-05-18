package main

import "sort"

/*
   动态规划转移方程：P(i,j)=P(i+1,j−1)∧(Si==Sj)
*/
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		var (
			palindrome string
			expand     bool
			extraIndex = i
		)
		palindrome = string(s[i])
		for j := i + 1; j < len(s) && j >= 0; j++ {
			if !expand && s[i] == s[j] {
				palindrome = s[i : j+1]
				continue
			} else {
				expand = true
			}
			if expand {
				// 中心开始扩展
				extraIndex--
				if extraIndex >= 0 && s[extraIndex] == s[j] {
					palindrome = s[extraIndex : j+1]
				} else {
					break
				}
			}
		}
		dp = append(dp, palindrome)
	}
	sort.Slice(dp, func(i, j int) bool {
		return len(dp[i]) > len(dp[j])
	})
	return dp[0]
}
