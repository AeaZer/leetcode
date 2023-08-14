package hash

/*给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。*/

// https://leetcode.cn/problems/group-anagrams/

import "sort"

func groupAnagrams(strings []string) [][]string {
	anagramsMap := map[string][]string{}
	for _, str := range strings {
		sb := []byte(str)
		sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
		sortedStr := string(sb)
		anagramsMap[sortedStr] = append(anagramsMap[sortedStr], str)
	}
	ans := make([][]string, 0, len(anagramsMap))
	for _, v := range anagramsMap {
		ans = append(ans, v)
	}
	return ans
}
