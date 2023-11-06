package everyday

import (
	"sort"
)

// 318. 最大单词长度乘积
// https://leetcode.cn/problems/maximum-product-of-word-lengths/description/?envType=daily-question&envId=2023-11-06

/*给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j]) 的最大值，
并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。*/

/*输入：words = ["abcw","baz","foo","bar","xtfn","abcdef"]
输出：16
解释：这两个单词为 "abcw", "xtfn"。*/

/*输入：words = ["a","aa","aaa","aaaa"]
输出：0
解释：不存在这样的两个单词。*/

// 贪心：长度倒序，乘积最大可以对乘积小的做剪枝
// 可以像官方题解那样使用位运算而不是 [26]bool，空间复杂度会更优秀，位运算一定程度上可以运算的更快，而不是遍历 26 个数组做比较
func maxProduct(words []string) int {
	wl := len(words)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	wordsCompresses := make([][26]bool, wl)
	for i := range words {
		for j := range words[i] {
			wordsCompresses[i][words[i][j]-'a'] = true
		}
	}
	var ans int
	for i := 0; i < wl; i++ {
		for j := i + 1; j < wl; j++ {
			l := len(words[i]) * len(words[j])
			if ans >= l {
				break
			}
			var isSame bool
			for z := 0; z < 26; z++ {
				if wordsCompresses[i][z] && wordsCompresses[j][z] {
					isSame = true
					break
				}
			}
			if !isSame {
				ans = l
			}
		}
	}
	return ans
}

// 使用位运算优化
func maxProduct2(words []string) int {
	wl := len(words)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	wordsCompresses := make([]int, wl)
	for i := range words {
		for j := range words[i] {
			wordsCompresses[i] |= 1 << (words[i][j] - 'a')
		}
	}
	var ans int
	for i := 0; i < wl; i++ {
		for j := i + 1; j < wl; j++ {
			l := len(words[i]) * len(words[j])
			if ans >= l {
				break
			}
			if wordsCompresses[i]&wordsCompresses[j] == 0 {
				ans = l
			}
		}
	}
	return ans
}
