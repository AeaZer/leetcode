package main

// 1525. 字符串的好分割数目
// https://leetcode.cn/problems/number-of-good-ways-to-split-a-string/description/

/*给你一个字符串 s ，一个分割被称为 「好分割」 当它满足：将 s 分割成 2 个字符串 p 和 q ，它们连接起来等于 s 且 p 和 q 中不同字符的数目相同。
请你返回 s 中好分割的数目。*/

func numSplits(s string) int {
	sl := len(s)
	left, right := make([]int, sl-1), make([]int, sl-1)
	left[0], right[0] = 1, 1
	leftBit, rightBit := uint8(0), uint8(0)
	for i, j := 1, sl-1; i < sl-1; {
		left[i] = left[i-1]
		if !hasSet(s[i], leftBit) {
			left[i]++
			leftBit |= 1 << (s[i] - 'a')
		}
		right[j] = right[j-1]
		if !hasSet(s[sl-i-1], rightBit) {
			right[j]++
			rightBit |= 1 << (s[sl-i-1] - 'a')
		}
		i++
		j++
	}
	var ans int
	for i := 0; i < sl; i++ {
		if left[i] == right[sl-i-1] {
			ans++
		}
	}
	return ans
}

func hasSet(b, num uint8) bool {
	return b-'a'&num != 0
}
