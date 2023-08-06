package doublepointer

/*给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：
如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？*/

// https://leetcode.cn/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
	// 如果 len(s) < len(t) 肯定不是子序列，直接返回空即可
	sl, tl := len(s), len(t)
	if sl > tl {
		return false
	}

	var tpFindIndex int
	for sp := 0; sp < sl; {
		spRune := s[sp]

		// 从上一个元素找到的位置下一个下标开始找当前元素
		tp := tpFindIndex
		for ; tp < tl; tp++ {
			if spRune == t[tp] {
				sp++
				tpFindIndex = tp + 1
				break
			}
		}
		// 找到最后都没找到，证明不是子序列
		if tp == tl {
			return false
		}

	}
	return true
}

// 进阶的场景，需要吧 T 序列的所有子序列查出来吧然后放在 map[subSeqStr]bool
