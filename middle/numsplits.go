package main

// 1525. 字符串的好分割数目
// https://leetcode.cn/problems/number-of-good-ways-to-split-a-string/description/

/*给你一个字符串 s ，一个分割被称为 「好分割」 当它满足：将 s 分割成 2 个字符串 p 和 q ，它们连接起来等于 s 且 p 和 q 中不同字符的数目相同。
请你返回 s 中好分割的数目。*/

func numSplits(s string) int {
	sl := len(s)
	left, right := make([]int, sl-1), make([]int, sl-1)
	left[0], right[0] = 1, 1
	leftBit, rightBit := 1<<(s[0]-'a'), 1<<(s[sl-1]-'a')
	for i := 1; i < sl-1; i++ {
		left[i] = left[i-1]
		lb := 1 << (s[i] - 'a')
		if !hasSet(lb, leftBit) {
			left[i]++
			leftBit |= lb
		}
		right[i] = right[i-1]
		rb := 1 << (s[sl-i-1] - 'a')
		if !hasSet(rb, rightBit) {
			right[i]++
			rightBit |= rb
		}
	}
	var ans int
	for i := 0; i < sl-1; i++ {
		if left[i] == right[sl-i-2] {
			ans++
		}
	}
	return ans
}

func hasSet(b, num int) bool {
	return b&num != 0
}
