package recall

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var (
		res       = make([][]int, 0)
		backtrack func([]int, int)
	)
	backtrack = func(path []int, index int) {
		if candidates[index] > target {
			return
		}
		if candidates[index]-target == 0 {
			comb := make([]int, len(path))
			copy(comb, path)
			res = append(res, comb)
			return
		}
		path = append(path, candidates[index])
		for i := 0; i < len(candidates); i++ {
			backtrack(path, i)
			path = path[:len(path)-1]
		}
	}
	backtrack([]int{}, 0)
	return res
}
