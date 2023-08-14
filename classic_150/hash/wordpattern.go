package hash

import "strings"

/*给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。*/

/*1 <= pattern.length <= 300
pattern 只包含小写英文字母
1 <= s.length <= 3000
s 只包含小写英文字母和 ' '
s 不包含 任何前导或尾随对空格
s 中每个单词都被 单个空格 分隔*/

// https://leetcode.cn/problems/word-pattern/

func wordPattern(pattern string, s string) bool {
	ses := strings.Split(s, " ")
	patternLen := len(pattern)
	if patternLen != len(ses) {
		return false
	}
	patternMap, sesMap := make(map[byte]string, 0), make(map[string]byte, 0)
	for i := 0; i < patternLen; i++ {
		sesValue, patternIndexValue := patternMap[pattern[i]], sesMap[ses[i]]
		if (sesValue != "" && sesValue != ses[i]) || (patternIndexValue != 0 && patternIndexValue != pattern[i]) {
			return false
		}
		patternMap[pattern[i]], sesMap[ses[i]] = ses[i], pattern[i]
	}
	return true
}
