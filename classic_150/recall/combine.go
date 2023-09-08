package recall

// 77. 组合
// https://leetcode.cn/problems/combinations/

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 你可以按 任何顺序 返回答案。

func combine(n int, k int) [][]int {
	nums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	var (
		res = make([][]int, 0)
		tmp = make([]int, 0, k)
	)
	var backtrack func(int)
	backtrack = func(index int) {
		tmpL := len(tmp)
		if index > n || tmpL+n-index+1 < k {
			return
		}
		if tmpL == k {
			comb := make([]int, k)
			copy(comb, tmp)
			res = append(res, comb)
			return
		}
		if index == n {
			return
		}
		tmp = append(tmp, nums[index])
		backtrack(index + 1)
		tmp = tmp[:len(tmp)-1]
		backtrack(index + 1)
	}
	backtrack(0)
	return res
}
