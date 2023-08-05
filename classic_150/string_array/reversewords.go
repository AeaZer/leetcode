package string_array

import "strings"

/*给你一个字符串 s ，请你反转字符串中 单词 的顺序。
单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。*/

// https://leetcode.cn/problems/reverse-words-in-a-string/

func reverseWords(s string) string {
	var (
		length        = len(s)
		words         = make([]string, 0)
		start         int
		currentLength int
	)

	for i := 0; i < length; i++ {
		if s[i] != ' ' {
			if currentLength == 0 {
				start = i
			}
			if i == length-1 {
				words = append(words, s[start:])
				break
			}
			currentLength++
			continue
		}
		// 之前是整个单词，现在把他塞进 words 数组里
		if currentLength > 0 {
			words = append(words, s[start:start+currentLength])
			currentLength = 0
		}
		// 之前就没有单词，是空格
	}

	var res strings.Builder
	for i := len(words) - 1; i >= 0; i-- {
		if res.Len() != 0 {
			res.WriteByte(' ')
		}
		res.WriteString(words[i])
	}

	return res.String()
}
