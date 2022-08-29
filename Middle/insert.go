package main

import "sort"

type twoDimensional [][2]int

func (t twoDimensional) Len() int {
	return len(t)
}
func (t twoDimensional) Less(i, j int) bool {
	return t[i][0] < t[j][0]
}
func (t twoDimensional) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func insert(intervals twoDimensional, newInterval []int) (res [][2]int) {
	sort.Sort(intervals)
	start := newInterval[0]
	end := newInterval[1]
	length := len(intervals)
	for i := 0; i < length; i++ {
		// [i,j] [i-1,j+1]
		if start > intervals[i][1] {
			res = append(res, intervals[i])
		}
		//[i,j] [i-1,j-1] => [i-1,j]
		if start <= intervals[i][0] && end <= intervals[i][0] {
			res = append(res, [2]int{start, intervals[i][1]})
		}
		//[i,j] [i-1,j+1] => [i-1,j+1]
		if start <= intervals[i][0] && end >= intervals[i][1] {
			res = append(res, [2]int{start, end})
		}
		//[i,j] [i+1,j-1] => [i,j]
		if start >= intervals[i][0] && end <= intervals[i][1] {
			res = append(res, intervals[i])
		}
		//[i,j] [i+1,j+1] => [i,j+1]
		if start >= intervals[i][0] && end <= intervals[i][1] {
			res = append(res, intervals[i])
		}
	}
	return
}
