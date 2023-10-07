package recall

import "strings"

// 22. 括号生成
// https://leetcode.cn/problems/generate-parentheses/description/?envType=study-plan-v2&envId=top-interview-150

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

const (
	parenthesisLeft = iota
	parenthesisRight
)

func buildTem(n int) [][]int {
	comb := make([][]int, 0)
	var callback func(int, int, []int)
	callback = func(l, r int, arr []int) {
		if r > l || r < 0 || l < 0 {
			return
		}
		if len(arr) == 2*n {
			tmp := make([]int, 2*n)
			copy(tmp, arr)
			comb = append(comb, tmp)
			return
		}
		if l < n {
			l++
			arr = append(arr, parenthesisLeft)
			callback(l, r, arr)
			l--
			arr = arr[:len(arr)-1]
		}
		if r < l {
			r++
			arr = append(arr, parenthesisRight)
			callback(l, r, arr)
			l--
			arr = arr[:len(arr)-1]
		}
	}
	callback(0, 0, []int{})
	return comb
}

func generateParenthesis(n int) []string {
	tem := buildTem(n)
	res := make([]string, 0, len(tem))
	for i := range tem {
		var b strings.Builder
		b.Grow(2 * n)
		for j := range tem[i] {
			if tem[i][j] == parenthesisLeft {
				b.WriteByte('(')
				continue
			}
			b.WriteByte(')')
		}
		res = append(res, b.String())
	}
	return res
}
