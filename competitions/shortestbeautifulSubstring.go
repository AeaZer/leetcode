package competitions

import "sort"

// 2904. 最短且字典序最小的美丽子字符串
// https://leetcode.cn/problems/shortest-and-lexicographically-smallest-beautiful-string/description/

/*给你一个二进制字符串 s 和一个正整数 k 。

如果 s 的某个子字符串中 1 的个数恰好等于 k ，则称这个子字符串是一个 美丽子字符串 。

令 len 等于 最短 美丽子字符串的长度。

返回长度等于 len 且字典序 最小 的美丽子字符串。如果 s 中不含美丽子字符串，则返回一个 空 字符串。
对于相同长度的两个字符串 a 和 b ，如果在 a 和 b 出现不同的第一个位置上，a 中该位置上的字符严格大于 b 中的对应字符，
则认为字符串 a 字典序 大于 字符串 b 。

例如，"abcd" 的字典序大于 "abcc" ，因为两个字符串出现不同的第一个位置对应第四个字符，而 d 大于 c 。*/

type beautifulPairs struct {
	fistIndex int
	lastIndex int
	str       string
}

// 思路：查找出所有合规的字符串，然后进行排序取第一位
func shortestBeautifulSubstring(s string, k int) string {
	pairs := make([]beautifulPairs, 0)
	sl := len(s)
	for i := 0; i < sl-k+1; i++ {
		c := k
		j := i
		for j < sl && c > 0 {
			if s[j] == '1' {
				c--
			}
			j++
		}
		if j <= sl && c == 0 {
			pairs = append(pairs, beautifulPairs{
				fistIndex: i,
				lastIndex: j,
				str:       s[i:j],
			})
		}
	}
	if len(pairs) == 0 {
		return ""
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].fistIndex-pairs[i].lastIndex != pairs[j].fistIndex-pairs[j].lastIndex {
			return pairs[i].fistIndex-pairs[i].lastIndex > pairs[j].fistIndex-pairs[j].lastIndex
		}
		return pairs[i].str < pairs[j].str
	})
	return s[pairs[0].fistIndex:pairs[0].lastIndex]
}
