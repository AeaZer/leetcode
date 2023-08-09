package sliding_window

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

// 思路：雙指針，如果 s[i:j] 是不重复的子串，那么 s[i+1, j] 也是不重复的，当含有重复的字符，j++ 一直到有重复的字符串为止，随后移动 i，重复这样做就可以得出正确的答案
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	nl := len(s)
	eMap := make(map[byte]bool, 0)
	maxLen := 0
	lp, rp := 0, 0
	for rp < nl {
		if !eMap[s[rp]] {
			eMap[s[rp]] = true
			rp++
			maxLen = max(maxLen, rp-lp)
			continue
		}
		delete(eMap, s[lp])
		lp++
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
