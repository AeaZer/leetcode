package everyday

// 187. 重复的DNA序列
// https://leetcode.cn/problems/repeated-dna-sequences/description/?envType=daily-question&envId=2023-11-05

/*DNA序列 由一系列核苷酸组成，缩写为 'A', 'C', 'G' 和 'T'.。

例如，"ACGAATTCCG" 是一个 DNA序列 。
在研究 DNA 时，识别 DNA 中的重复序列非常有用。

给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的 长度为 10 的序列(子字符串)。你可以按 任意顺序 返回答案。
*/

/*输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出：["AAAAACCCCC","CCCCCAAAAA"]*/

func findRepeatedDnaSequences(s string) []string {
	sm := make(map[string]int, 0)
	sl := len(s)
	for i := 0; i < sl-9; i++ {
		sm[s[i:i+10]]++
	}
	ans := make([]string, 0, len(sm))
	for key, count := range sm {
		if count > 1 {
			ans = append(ans, key)
		}
	}
	return ans
}
