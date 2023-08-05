package string_array

import "sort"

/*编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。*/

// https://leetcode.cn/problems/longest-common-prefix/

// 木桶效应，公共最长的前缀将最短的字符串作为依据
func longestCommonPrefix(strings []string) string {
	if len(strings) == 1 {
		return strings[0]
	}

	sort.Slice(strings, func(i, j int) bool {
		return len(strings[i]) < len(strings[j])
	})

	var (
		minStrBytes   = []byte(strings[0])
		haveDiff      bool
		equalMaxIndex int
	)

	for equalMaxIndex < len(minStrBytes) {
		for i := 1; i < len(strings); i++ {
			if strings[i][equalMaxIndex] != minStrBytes[equalMaxIndex] {
				haveDiff = true
				break
			}
		}
		if haveDiff {
			break
		}
		equalMaxIndex++
	}

	return strings[0][:equalMaxIndex]
}
