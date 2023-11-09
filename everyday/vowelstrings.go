package everyday

// 2586. 统计范围内的元音字符串数
// https://leetcode.cn/problems/count-the-number-of-vowel-strings-in-range/description/?envType=daily-question&envId=2023-11-07

/*给你一个下标从 0 开始的字符串数组 words 和两个整数：left 和 right 。
如果字符串以元音字母开头并以元音字母结尾，那么该字符串就是一个 元音字符串 ，
其中元音字母是 'a'、'e'、'i'、'o'、'u' 。
返回 words[i] 是元音字符串的数目，其中 i 在闭区间 [left, right] 内。*/

func vowelStrings(words []string, left int, right int) int {
	var ans int
	for i := left; i <= right; i++ {
		if isVowelCharacter(words[i][0]) &&
			isVowelCharacter(words[i][len(words[i])-1]) {
			ans++
		}
	}
	return ans
}

func isVowelCharacter(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}
