package interval

import "sort"

func merge(intervals [][]int) [][]int {
	iL := len(intervals)
	if iL == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	slow, fast := 0, 0
	for fast < iL {
		fast++
		if fast == iL || intervals[fast-1][1] < intervals[fast][0] {
			res = append(res, []int{intervals[slow][0], intervals[fast-1][1]})
			slow = fast
			continue
		}
		if intervals[fast-1][1] > intervals[fast][1] {
			intervals[fast][1] = intervals[fast-1][1]
		}
	}
	return res
}
