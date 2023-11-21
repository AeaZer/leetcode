package huawei

import (
	"strings"
)

/*输入单行英文句子，里面包含英文字母，空格以及,.?
三种标点符号，请将句子内每个单词进行倒序，并输出倒序后的语句
输入描述
输入字符串 S，S 的长度 1≤N≤100
输出描述
输出逆序后的字符串*/

func wordFlashback(str string) string {
	sl := len(str)
	var b strings.Builder
	b.Grow(sl)
	for i := 0; i < sl; {
		if isBreakChar(str[i]) {
			b.WriteByte(str[i])
			i++
			continue
		}
		slow, fast := i, i
		for fast < sl && !isBreakChar(str[fast]) {
			fast++
		}
		for j := fast - 1; j >= slow; j-- {
			b.WriteByte(str[j])
		}
		i = fast
	}
	return b.String()
}

func isBreakChar(ch byte) bool {
	if ch == ' ' || ch == '.' || ch == ',' || ch == '?' {
		return true
	}
	return false
}
