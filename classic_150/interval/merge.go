package interval

import "sort"

/*以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。*/

// https://leetcode.cn/problems/merge-intervals/

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
