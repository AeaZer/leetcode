package binarysearch

import "testing"

func TestSearchRange(t *testing.T) {
	res := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	t.Log(res)
}
