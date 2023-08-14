package hash

import "sort"

/*给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。*/

// https://leetcode.cn/problems/valid-anagram/

/*1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// key: 某个字符，value：这个字符出现的次数
	sm := make(map[byte]uint16, 26)
	for i := range s {
		sm[s[i]]++
	}
	for i := range t {
		// 不存在某个字符或某个字符不够用了
		if v, ok := sm[t[i]]; !ok || v == 0 {
			return false
		}
		sm[t[i]]--
	}
	return true
}

// 利用了字符串底层是字节数组进行两次类型转换以及排序
func isAnagram2(s string, t string) bool {
	sb, tb := []byte(s), []byte(t)
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	sort.Slice(tb, func(i, j int) bool {
		return tb[i] < tb[j]
	})
	return string(sb) == string(tb)
}
