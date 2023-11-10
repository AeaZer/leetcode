package main

import "strings"

// 1768. 交替合并字符串
// https://leetcode.cn/problems/merge-strings-alternately/description/

/*
给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。
如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。
返回 合并后的字符串 。
*/
func mergeAlternately(word1 string, word2 string) string {
	n, m := len(word1), len(word2)
	var b strings.Builder
	b.Grow(n + m)
	var index int
	for ; index < n && index < m; index++ {
		b.WriteByte(word1[index])
		b.WriteByte(word2[index])
	}
	if index < n {
		b.WriteString(word1[index:])
	}
	if index < m {
		b.WriteString(word2[index:])
	}
	return b.String()
}
