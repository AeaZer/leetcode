package main

// 56. 合并区间
// 一维数组排序
func merge(intervals [][]int) [][]int {
	// 排序
	size := len(intervals)
	sorted := quickSort(intervals, 0, size-1)

	res := make([][]int, 0, size)
	for i := 0; i < size; i++ {
		start := sorted[i]
		temp := start
		for j := i + 1; j < size; j++ {
			end := sorted[j]

			if start[1] <= end[0] {
				start = end
				continue
			}
			res = append(res, append([]int(nil), []int{temp[0], start[1]}...))
			i = j
			break
		}
	}

	return res
}
