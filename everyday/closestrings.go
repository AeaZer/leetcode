package everyday

import "sort"

// 1657. 确定两个字符串是否接近
// https://leetcode.cn/problems/determine-if-two-strings-are-close/description/?envType=daily-question&envId=2023-11-30

/*如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：

操作 1：交换任意两个 现有 字符。
例如，abcde -> aecdb
操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
你可以根据需要对任意一个字符串多次使用这两种操作。

给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。
*/

// 思路：统计字符出现的次数，既然字符串可以相互互换，那么其实对于两个字符串来说
// 字符串相近只需要排除特性情况就好
// 字符的每次出现转换为另一个 现有 字符这个就就是其中的一条特殊情况，排除掉就好

func closeStrings(word1 string, word2 string) bool {
	l1, l2 := len(word1), len(word2)
	if l1 != l2 {
		return false
	}
	ws1, ws2 := [26]int{}, [26]int{}
	for i := 0; i < l1; i++ {
		ws1[word1[i]-'a']++
		ws2[word2[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if ws1[i] > 0 && ws2[i] == 0 || ws1[i] == 0 && ws2[i] > 0 {
			return false
		}
	}
	sort.Ints(ws1[:])
	sort.Ints(ws2[:])
	for i := 0; i < 26; i++ {
		if ws1[i] != ws2[i] {
			return false
		}
	}
	return true
}