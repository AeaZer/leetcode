package everyday

// 2609. 最长平衡子字符串
// https://leetcode.cn/problems/find-the-longest-balanced-substring-of-a-binary-string/description/?envType=daily-question&envId=2023-11-08

/*给你一个仅由 0 和 1 组成的二进制字符串 s 。
如果子字符串中 所有的 0 都在 1 之前 且其中 0 的数量等于 1 的数量，则认为 s 的这个子字符串是平衡子字符串。请注意，空子字符串也视作平衡子字符串。
返回  s 中最长的平衡子字符串长度。
子字符串是字符串中的一个连续字符序列*/

/*输入：s = "01000111"
输出：6
解释：最长的平衡子字符串是 "000111" ，长度为 6 。*/

func findTheLongestBalancedSubstring(s string) int {
	sl := len(s)
	var ans int
	for i := 0; i < sl; {
		var zc int
		for ; i < sl && s[i] == '0'; i++ {
			zc++
		}
		var oc int
		for ; i < sl && s[i] == '1'; i++ {
			oc++
		}
		// 木桶效应
		ans = max(min(zc, oc), ans)
	}
	return ans * 2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
