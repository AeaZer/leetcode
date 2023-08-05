package doublepointer

import (
	"strings"
	"unicode"
)

/*如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
字母和数字都属于字母数字字符。
给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。*/

// https://leetcode.cn/problems/valid-palindrome/

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	for lp, rp := 0, len(s)-1; lp < rp; {
		if s[lp] == s[rp] {
			lp++
			rp--
			continue
		}
		if !unicode.IsDigit(rune(s[lp])) && !unicode.IsLetter(rune(s[lp])) {
			lp++
			continue
		}
		if !unicode.IsDigit(rune(s[rp])) && !unicode.IsLetter(rune(s[rp])) {
			rp--
			continue
		}
		return false
	}
	return true
}
