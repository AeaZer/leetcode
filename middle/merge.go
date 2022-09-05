package main

// 56. 合并区间
func merge(intervals [][]int) [][]int {
	// 排序
	size := len(intervals)
	sorted := quickSort(intervals, 0, size-1)

	res := make([][]int, 0, size)
	for i := 0; i < size; i++ {
		disable := false
		start := sorted[i]
		// 存入当前值
		temp := start
		j := i + 1
		resSize := len(res)
		if resSize > 0 && res[resSize-1][1] > start[1] {
			continue
		}
		for !disable && j < size {
			//解决[[1,3],[4,5]] ; [[2,3],[2,2]] 的问题
			if start[1] >= sorted[j][0] && start[1] < sorted[j][1] {
				start = sorted[j]
				j++
				continue
			}
			if start[1] >= sorted[j][1] {
				j++
			}
			disable = true
		}
		res = append(res, append([]int(nil), []int{temp[0], start[1]}...))
		i = j - 1
	}

	return res
}
