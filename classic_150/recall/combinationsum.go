package recall

import "sort"

func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
	length := len(candidates)
	res := make([][]int, 0)
	var callback func([]int, int, int)
	callback = func(arr []int, sum, index int) {
		for i := index; i < length; i++ {
			arr = append(arr, candidates[i])
			sum += candidates[i]
			if sum == target {
				comb := make([]int, len(arr))
				copy(comb, arr)
				res = append(res, comb)
			}
			if sum > target {
				return
			}
			callback(arr, sum, i)
			arr = arr[:len(arr)-1]
			sum -= candidates[i]
		}
	}
	callback([]int{}, 0, 0)
	return res
}
