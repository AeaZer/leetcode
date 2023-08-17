package interval

/*给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。*/

// https://leetcode.cn/problems/insert-interval/

func insert(intervals [][]int, newInterval []int) [][]int {
	iL := len(intervals)
	if iL == 0 {
		return [][]int{newInterval}
	}
	newIntervals := make([][]int, iL+1)
	newIntervals[0] = newInterval
	copy(newIntervals[1:], intervals)
	return merge(newIntervals)
}
