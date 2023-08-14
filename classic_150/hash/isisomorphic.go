package hash

/*给定两个字符串 s 和 t ，判断它们是否是同构的。
如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。*/

/*1 <= s.length <= 5 * 104
t.length == s.length
s 和 t 由任意有效的 ASCII 字符组成*/

// https://leetcode.cn/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	sl, tl := len(s), len(t)
	if sl != tl {
		panic("unexpect s,t length")
	}
	sm, tm := make(map[byte]byte, 0), make(map[byte]byte, 0)
	for i := 0; i < sl; i++ {
		x, y := sm[s[i]], tm[t[i]]
		if (x > 0 && x != t[i]) || (y > 0 && y != s[i]) {
			return false
		}
		sm[s[i]], tm[t[i]] = t[i], s[i]
	}
	return true
}
