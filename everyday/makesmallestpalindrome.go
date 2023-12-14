package everyday

// 2697. 字典序最小回文串
// https://leetcode.cn/problems/lexicographically-smallest-palindrome/description/

func minByte (a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

func makeSmallestPalindrome(s string) string {
	var (
		left = 0
		right = len(s)-1
	)
	sb := []byte(s)
	for left < right {
		sb[left] = minByte(sb[left], sb[right])
		sb[right] = sb[left]
		left++
		right--
	}
	return string(sb)
}